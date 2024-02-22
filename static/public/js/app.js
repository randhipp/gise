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
		child.innerText = images[index].url + " " + images[index].hue;

		container.appendChild(child);
	}
};

$('#add_image').addEventListener('click', async (e) => {
	e.preventDefault();
	const image = $('#image').value;
	const hue = $('#hue').value;

	if (!image) return;
	if (!hue) return;

	const form = new FormData();
	form.append('image', image);
	form.append('hue', hue);

	const response = await fetch(API_ENDPOINT, {
		method: 'POST',
		body: form,
	});

	for (var pair of form.entries()) {
		console.log(pair[0]+ ', ' + pair[1]); 
	}
	const data = await response.json();

	const child = document.createElement('li');
	child.className = 'list-group-item';
	child.innerText = data.image.url + " " + data.image.hue;

	container.insertBefore(child, container.firstChild);

	$('#image').value = '';
	$('#hue').value = '0';
});

document.addEventListener('DOMContentLoaded', listImages);
