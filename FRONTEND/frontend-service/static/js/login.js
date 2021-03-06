const username1 = $('#username_field')
const password = $('#password_field')

function login() {

    var obj = {
        username : username1.val(),
        password : password.val()
    }
    
    $.ajax({
        type:'POST',
        crossDomain: true,
        url: AUTH_SERVICE_URL + '/login',
        data : JSON.stringify(obj),
        contentType : 'application/json',
        //dataType: 'JSON',
        success : function(data) {
            console.log(data.token)
            console.log('your token is' + data.token)
            localStorage.setItem('myToken', data.token);
            //alert("Succesfully logged in")
            redirectMe()
            //alert(data.token)
        },
        error : function(xhr, status, data) {
            console.log(xhr)
            console.log('Error in login');
            //alert('Wrong credentials!')
            showLoginError('Bad credentials', 'Please recheck your credentials.')
        }
    })
}

function redirectMe() {
    whoamiredirect()
}

function logout(){
    localStorage.removeItem('myToken');
    window.location.href = 'login.html';
}

function whoamiredirect() {

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
            if (data.role == 'admin') {
                window.location.href = 'admin.html'
                return
            }
            console.log('redirecting user to feed page')
            window.location.href = 'feed.html'

        },
        error: function () {

            alert('Could not redirect user')
        }
    })
}