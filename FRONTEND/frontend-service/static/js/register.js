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

    if(validateInputs() == false) {
        return;
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

            clearRegFields()
            redirectToLoginPage();
            alert('Successfull registration, please log in')

        },
        error : function(xhr, status, data) {
            console.log(xhr.status);
      
            if(xhr.status == 417) {
                alert("User exists with this username or email!")

            } else {
                // to be checked in the future why after successfull reg drops error...
                alert('Successfull registration, please log in')
                clearRegFields();
                redirectToLogin();

            }

        }
    })
}

function clearRegFields() {
    mail.val('')
    username.val('')
    name.val('')
    lastname.val('')
    password.val('')
}

function redirectToLogin() {
    console.log('redirecting user to login page after registration')
    window.location.href = 'login.html'
}

function validateInputs() {
    if(mail.val() =="") {
        alert("Mail field can't be null");
        return false
    }

    if(username.val() =="") {
        alert("Username field can't be null");
        return false
    }

    if(name.val() =="") {
        alert("Name field can't be null");
        return false
    }

    if(lastname.val() =="") {
        alert("Lastname field can't be null");
        return false
    }

    if(password.val() =="") {
        alert("Password field can't be null");
        return false
    }

    var pattern = username.attr("pattern");
    var re = new RegExp(pattern);

    if(!re.test(username.val())) {
        alert("Invalid character in username field // must be between 3-20 characters");
        return false;
    }

    pattern =name.attr("pattern");
    re = new RegExp(pattern);

    if(!re.test(name.val())) {
        alert('Invalid character in name field');
        return false;
    }

    if(!re.test(lastname.val())) {
        alert('Invalid character in lastname field');
        return false;
    }

    return true

}