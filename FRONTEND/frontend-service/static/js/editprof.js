var currentprofileObj;
// edit prof data
const name = $("#nameid");
const lastname = $("#lastnameid");
const username = $("#usernameid");

const email = $("#emailid");
const phone = $("#phoneid");
const gender = $("#genderid");
const dateofbirth = $("#birthid");
const bio = $("#boiid");
const website = $("#websiteid");

//privacy 
const profprivacy = $("#profprivacy_id");
const dmprivacy = $("#dmprivacy_id");
const tagprivacy = $("#tagprivacy_id");

//notif
const follownotif = $("#follownotif_id");
const dmnotif = $("#dmnotif_id");
const tagnotif = $("#tagnotif_id");


function getMyDatas() {
    // getdata from user
    $.ajax({
        type:'GET',
        crossDomain: true,
        url: AUTH_SERVICE_URL + '/getdata',
        contentType : 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            //console.log(data)
            fillProfileDataFromUser(data)
        },
        error : function(xhr, status, data) {
            //console.log(xhr)
            //console.log('Cant get profile data');
        }
    })

    // getdata from profile
    $.ajax({
        type:'GET',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/getdata',
        contentType : 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success : function(data) {
            //console.log(data)
            fillProfileDataFromProfile(data)
            fillPrivacySettings(data)
            fillNotifSettings(data)

            currentprofileObj = data
            checkIfUserIsMutedOrBlocked()
            fillUnMuteTable(currentprofileObj)
            fillUnBlockTable(currentprofileObj)
        },
        error : function(xhr, status, data) {
            //console.log(xhr)
            //console.log('Cant get user data');
        }
    })

   
}

function fillProfileDataFromUser(data) {
    // name, lastname, username

    email.val(data.email)
    name.val(data.name)
    lastname.val(data.lastname)
    username.val(data.username)
}

function fillProfileDataFromProfile(data) {

    phone.val(data.phone)

    //console.log(data.gender)
    if(data.gender == 0) {
        //console.log("GENDER = 0")
        $("#genderid").val('1');

    } else {
        //console.log("GENDER = 1")
        $("#genderid").val('2');
    }
    //gender.val(data.gender)

    // ?????  why undefined 
    ////console.log(data.date_of_birth)
    
    if(data.date_of_birth == "" || data.date_of_birth == "undefined" || data.date_of_birth == null) {
        dateofbirth.val("")
    } else {
        dateofbirth.val(data.date_of_birth.toString())
    }

    bio.val(data.biography)
    website.val(data.website)
}

function fillPrivacySettings(data) {
    //console.log("PRIVACY")
    //console.log(data.privacy_setting)
    
    if(data.privacy_setting.is_public == true) {
        $(function(){ $('#profprivacy_id').bootstrapToggle('on') });
    } else {
        $(function(){ $('#profprivacy_id').bootstrapToggle('off') });
    }

    if(data.privacy_setting.is_inbox_open == true) {
        $(function(){ $('#dmprivacy_id').bootstrapToggle('on') });
    } else {
        $(function(){ $('#dmprivacy_id').bootstrapToggle('off') });
    }

    if(data.privacy_setting.is_tagging_allowed == true) {
        $(function(){ $('#tagprivacy_id').bootstrapToggle('on') });
    } else {  
        $(function(){ $('#tagprivacy_id').bootstrapToggle('off') });
    }

}

function fillNotifSettings(data) {
    //console.log("NOTIF")
    //console.log(data.notification_setting)

    if(data.notification_setting.show_follow_notification == true) {
        $(function(){ $('#follownotif_id').bootstrapToggle('on') });
    } else {
        $(function(){ $('#follownotif_id').bootstrapToggle('off') });
    }

    if(data.notification_setting.show_dm_notification == true) {
        $(function(){ $('#dmnotif_id').bootstrapToggle('on') });
    } else {
        $(function(){ $('#dmnotif_id').bootstrapToggle('off') });
    }

    if(data.notification_setting.show_tagged_notification == true) {
        $(function(){ $('#tagnotif_id').bootstrapToggle('on') });
    } else {  
        $(function(){ $('#tagnotif_id').bootstrapToggle('off') });
    }
}

function fillUnMuteTable(currentprofileObj) {
    var table = document.getElementById('muteTableBody')

    table.innerHTML = ''
    for(let i = 0; i < currentprofileObj.graylist.length; i++ ) {

        var usernameToUnmute = currentprofileObj.graylist[i].username

        var row = `<tr>
                        <td>${currentprofileObj.graylist[i].username}</td>
                        <td><button class="btn btn-dark insidetablebtn" id="`+ usernameToUnmute +`"  onclick='unmuteUserWithParam(this.id)'>Unmute</button></td>
                    </tr>`

        table.innerHTML += row
    }

}

function fillUnBlockTable(currentprofileObj) {
    var table2 = document.getElementById('blockTableBody')
    table2.innerHTML = ''

    for(let i = 0; i < currentprofileObj.blacklist.length; i++ ) {

        var usernameToUnblock = currentprofileObj.blacklist[i].username

        var row = `<tr>
                        <td>${currentprofileObj.blacklist[i].username}</td>
                        <td><button class="btn btn-dark insidetablebtn" id="`+ usernameToUnblock +`"  onclick='unblockUserWithParam(this.id)'>Unblock</button></td>
                    </tr>`

        table2.innerHTML += row
    }

}