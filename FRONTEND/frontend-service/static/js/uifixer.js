$(document).ready(function() {
	redirectCheck()
})


function hideNavButtons() {
	$('#homenav').hide()
	$('#reqsnav').hide()
	$('#profnav').hide()
	$('#logoutnav').hide()
	$('#loginnav').html('Login')
}

function redirectCheck() {
	$.ajax({
        type: 'GET',
        crossDomain: true,
        url: AUTH_SERVICE_URL + '/whoami',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function (data) {
			

        },
        error: function () {
			//alert("NISI ULOGOVAN")
			if (window.location.href.includes("profile")) {
				//alert("NISI ULOGOVAN")
				$('#homenav').hide()
				$('#reqsnav').hide()
				$('#profnav').hide()
				$('#logoutnav').hide()
				$('#postnotifsnav').hide()
				$('#loginnavtext').html("Login")
			}
        }
    })
	

}
