import { extractAttributes } from "./tools/attribute.js";
import { addToElemValue } from "./tools/math.js";

interface LikeDislikePostRequestAjax {
	post_id: number;
	user_id: number;
}

interface LikeDislikePostAttributeMap {
	post_id: string
	user_id: string
}

const	LIKE_DISLIKE_POST_ATTRIBUTE_MAP: LikeDislikePostAttributeMap = {
	post_id: 'data-post-id',
	user_id: 'data-user-id'
}

function sendLikeDislikeRequest(button: HTMLElement, action: string): number {
	let		data:		LikeDislikePostAttributeMap | null;
	let		request:	LikeDislikePostRequestAjax | null;

	data = extractAttributes<LikeDislikePostAttributeMap>(
		button,
		LIKE_DISLIKE_POST_ATTRIBUTE_MAP
	);
	if (!data)
		return 1
	request = {
		post_id: parseInt(data.post_id, 10),
		user_id: parseInt(data.user_id, 10)
	}
	fetch(action, {
		method: 'POST',
		headers: {
			'Content-Type': 'application/json',
		},
		body: JSON.stringify(request)
	})
	.catch((error: Error) => {
		console.error('Error:', error);
		alert('Something went wrong, please try again.');
	});
	return 0
}

function addToLikeDislikeButtonValue(button: HTMLButtonElement, nb: number) {
	let		buttonCount:	HTMLElement | null;

	buttonCount = button.querySelector('.like-dislike-count')
	if (!buttonCount) {
		console.error("Element with class like-dislike-count not found")
		return
	}
	addToElemValue(buttonCount, nb)
}

function handleLikeDislikeButton(event: Event, action: string,
									oppositeButton: HTMLButtonElement): void {
	const	button = event.currentTarget as HTMLButtonElement | null;

	if (!button)
		return
	if (sendLikeDislikeRequest(button, action) == 1)
		return
	if (!button.classList.contains('active')) {
		addToLikeDislikeButtonValue(button, 1);
	} else {
		addToLikeDislikeButtonValue(button, -1);
	}
	button.classList.toggle('active');
	if (oppositeButton.classList.contains('active')) {
		addToLikeDislikeButtonValue(oppositeButton, -1);
		oppositeButton.classList.remove('active');
	}
}

document.addEventListener('DOMContentLoaded', () => {
	const likeButton = document.getElementById(
							'likeButton') as HTMLButtonElement;
	const dislikeButton = document.getElementById(
							'dislikeButton') as HTMLButtonElement;

	likeButton.addEventListener('click',
		(event) => handleLikeDislikeButton(event, "/post/like", dislikeButton)
	);
	dislikeButton.addEventListener('click',
		(event) => handleLikeDislikeButton(event, "/post/dislike", likeButton)
	);
});
