const modal = $('#new-post-modal')
var submit_endpoint = '/post/upload'
var sent = false
function showModalPost() {
	modal.modal('show')
}

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

$('#select_type').change(function() {


	if ($("#select_type").prop('selectedIndex') == 0) {
		console.log('Post')
		submit_endpoint = '/post/upload'
	}
	else if ($("#select_type").prop('selectedIndex') == 2) {
		console.log('Story')
		submit_endpoint = '/story/upload'
	}
	else {
		console.log('Album')
		submit_endpoint = '/post/upload'
	}
})



function changeActionURL() {	
	
	$('#post_form').attr('action', CONTENT_SERVICE_URL + submit_endpoint); 

	//e.preventDefault(); // avoid to execute the actual submit of the form.

	var datax = new FormData($("#post_form")[0]);
	
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
           success: function(data)
           {
               alert(data); // show response from the php script.
           }
         });


	
}