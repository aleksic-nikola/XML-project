const username = $('#username_field')
const password = $('#password_field')

function login() {

    var obj = {
        username : username.val(),
        password : password.val()
    }
    
    $.ajax({
        type:'POST',
        crossDomain: true,
        url: 'http://localhost:9090/login',
        data : JSON.stringify(obj),
        //headers: { 'Content-Type': 'application/json' },
        contentType : 'application/json',
        success : function(data) {
            console.log(data)
        },
        error : function(xhr, status, data) {
            console.log(xhr)
            console.log('Cant get msgs');
        }
    })
}