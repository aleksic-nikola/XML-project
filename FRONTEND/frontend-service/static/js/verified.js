const verification_category_select = $('#verification_category_select')
const verification_form = $('#verification_form')
verification_form.submit(function () {
	sendVerificationRequest();
	return false;
       });

function sendVerificationRequest() {

	if ($('#verification_category').val().trim() == '') {
		alert('You must choose a category')
		return 
	}

	if ($('#verification_name').val().trim() == '') {
		alert('You must input your real name')
		return 
	}

	if ($('#verification_lastname').val().trim() == '') {
		alert('You must input your real last name')
		return 
	}

	verification_form.attr('action', REQUEST_SERVICE_URL + '/verificationReqs/create'); 

	//e.preventDefault(); // avoid to execute the actual submit of the form.

	var datax = new FormData($("#verification_form")[0]);


	
	//var dataz = form.serialize()
	//console.log(dataz)
    	$.ajax({
           type: "POST",
           url: REQUEST_SERVICE_URL + '/verificationReqs/create',
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

verification_category_select.change(function() {

	var category = verification_category_select.children(":selected").attr("id");
	console.log(category)
	$('#verification_category').val(category)

})