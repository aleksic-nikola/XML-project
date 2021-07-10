const modal = $('#new-post-modal')
const location_field = $('#post_title')
const tags_field = $('#description_part')
const album_select = $('#album_select')
const closeFriend_select = $("#closeFriend_select")
var submit_endpoint = '/post/upload'
var sent = false
function showModalPost() {
	modal.modal('show')
}

$(document).ready(function () {

	album_select.hide()
	closeFriend_select.hide()
})

function postSelected() {
	/*
	event.preventDefault()
	var form = new FormData($("#uploadForm")[0])
	console.log(form)
	return
	ajax({
		type: 'POST', // For jQuery < 1.9
		url: CONTENT_SERVICE_URL + /post/upload,
		data: form,
		cache: false,
		contentType: false,
		processData: false,
		//method: 'POST',
		success: function(data){
			alert(data);
		}
	});
	*/
}

$('#select_type').change(function () {


	if ($("#select_type").prop('selectedIndex') == 0) {
		console.log('Post')
		submit_endpoint = '/post/upload'
		tags_field.show()
		location_field.show()
		album_select.hide()
		$("#closeFriend_select").hide()
	}
	else if ($("#select_type").prop('selectedIndex') == 2) {
		console.log('Story')
		submit_endpoint = '/story/upload'
		tags_field.hide()
		location_field.hide()
		album_select.hide()
		$("#closeFriend_select").show()

	}
	else {
		console.log('Album')
		tags_field.show()
		location_field.show()
		album_select.show()
		//submit_endpoint = '/post/upload'

	}
})

album_select.change(function () {

	var index = album_select.prop('selectedIndex')

	if (index == 0) {
		submit_endpoint = ''
	}
	else if (index == 1) {
		submit_endpoint = '/post/upload'
		tags_field.show()
		location_field.show()
		tags_field.show()
		location_field.show()
		$("#closeFriend_select").hide()

	}
	else {
		submit_endpoint = '/story/upload'
		tags_field.hide()
		location_field.hide()
		tags_field.hide()
		location_field.hide()
		$("#closeFriend_select").show()

	}

})



function changeActionURL() {

	$('#post_form').attr('action', CONTENT_SERVICE_URL + submit_endpoint);

	//e.preventDefault(); // avoid to execute the actual submit of the form.

	// only if on story 
	if ($("#select_type").prop('selectedIndex') == 2) {
		if ($("#closeFriend_select").prop('selectedIndex') == 0) {
			alert("Select if is story for close friend")
			return
		}
		else if ($("#closeFriend_select").prop('selectedIndex') == 1) { //true
			$("#closeFriends_input").val("TRUE")
		}
		else {
			$("#closeFriends_input").val("FALSE")
	
		}
	}
	// only if on album and if on story
	if ($("#select_type").prop('selectedIndex') == 1 && album_select.prop('selectedIndex') == 2) {
		if ($("#closeFriend_select").prop('selectedIndex') == 0) {
			alert("Select if is story for close friend")
			return
		}
		else if ($("#closeFriend_select").prop('selectedIndex') == 1) { //true
			$("#closeFriends_input").val("TRUE")
		}
		else {
			$("#closeFriends_input").val("FALSE")
	
		}
	}

	var datax = new FormData($("#post_form")[0]);
	//multiplefiles
	var inp = document.getElementById('multiplefiles');

	var type_index = $("#select_type").prop('selectedIndex')

	// post 
	if (type_index == 0) {
		console.log('trying to send a post')
		if (inp.files.length != 1) {
			alert('You can only choose 1 image/video for post')
			return
		}
		console.log(submit_endpoint.toUpperCase())
		// story
	} else if (type_index == 2) {

		console.log('trying to send a story')
		if (inp.files.length != 1) {
			alert('You can only choose 1 image/video for story')
			return
		}
		console.log(submit_endpoint.toUpperCase())
	}
	// album
	else {
		console.log('trying to send an album')
		var album_index = album_select.prop('selectedIndex')

		if (album_index == 0) {
			alert('Please choose what you want to post your album as!')
			return
		}
		else if (album_index == 1) {
			// post
			if (inp.files.length < 2) {
				alert('You have to select multiple images/videos to post an album')
				return
			}
			submit_endpoint = '/post/upload'
			console.log(submit_endpoint.toUpperCase())
		} else {
			// story 
			if (inp.files.length < 2) {
				alert('You have to select multiple images/videos to post an album')
				return
			}
			submit_endpoint = '/story/upload'
			console.log(submit_endpoint.toUpperCase())
		}
	}

	// for (var i = 0; i < inp.files.length; ++i) {
	// 	var name = inp.files.item(i).name;
	// 	alert("here is a file name: " + name);
	// }

	//var dataz = form.serialize()
	//console.log(dataz)
	$.ajax({
		type: "POST",
		url: CONTENT_SERVICE_URL + submit_endpoint,
		//contentType: 'multipart/form-data',
		data: datax, // serializes the form's elements.
		beforeSend: function (xhr) {
			xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		processData: false,
		contentType: false,
		success: function (data) {
			alert(data); // show response from the php script.
		}
	});



}