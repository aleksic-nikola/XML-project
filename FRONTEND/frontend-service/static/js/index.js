$(document).ready(function() {
    checkUser();
})

function checkUser() {
    $.ajax({
        type:'GET',
        crossDomain: true,
        url: AUTH_SERVICE_URL + '/whoami',
        contentType : 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            console.log(data)
            redirectToFeedPage()
        },
        error : function(xhr) {
            console.log(xhr)
            redirectToLoginPage()
            console.log('ERROR IN WHOAMI');
        }
    })

}

function redirectToLoginPage() {
    console.log('redirecting user to login page')
    window.location.href = 'login.html'
}

function redirectToFeedPage() {
    console.log('redirecting user to feed page')
    window.location.href = 'feed.html'
}

