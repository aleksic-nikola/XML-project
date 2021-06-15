const mail = $("#mail_input");
const username = $("#username_input");
const name = $("#name_input");
const lastname = $("#lastname_input");
const password = $("#password_input");

function register() {

    var obj = {
        "email" : mail.val(),
        "username" : username.val(),
        "name" : name.val(),
        "lastname" : lastname.val(),
        "password" : password.val(),
        "role" : "user"
    }

    $.ajax({
        type:'POST',
        crossDomain: true,
        url: 'http://localhost:9090/register',
        //headers: { 'Content-Type': 'application/json' },
        contentType : 'application/json',
        data: JSON.stringify(obj),
        success : function(data) {
            console.log(data)
        },
        error : function(xhr, status, data) {
            console.log(xhr)
            console.log('Cant get msgs');
        }
    })
}