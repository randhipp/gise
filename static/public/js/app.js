const $ = (selector) => document.querySelector(selector);
const container = $('#images');
const API_ENDPOINT = '/api/v1/images';

const listImages = async () => {
	const response = await fetch(API_ENDPOINT);
	const data = await response.json();
	const images = data.images.reverse();

	for (let index = 0; index < images.length; index++) {
		const child = document.createElement('li');
		child.className = 'list-group-item';
		child.innerText = images[index].name;

		container.appendChild(child);
	}
};

$('#add_image').addEventListener('click', async (e) => {
	e.preventDefault();
	const image = $('#image').value;

	if (!image) return;

	const form = new FormData();
	form.append('image', image);

	const response = await fetch(API_ENDPOINT, {
		method: 'POST',
		body: form,
	});

	const data = await response.json();

	const child = document.createElement('li');
	child.className = 'list-group-item';
	child.innerText = data.image.name;

	container.insertBefore(child, container.firstChild);

	$('#image').value = '';
});

document.addEventListener('DOMContentLoaded', listImages);
