/*
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
*/

function saveNewData() {

	console.log(whoCanISee)

	if (whoCanISee == 'profile') {
		saveProfileData()
		return
	}

	if (whoCanISee == 'privacy') {
		savePrivacyData()
		return
	}

	if (whoCanISee == 'notification') {
		saveNotificationData()
		return
	}

	if (whoCanISee == 'muteblock') {
		saveMuteBlockData()
		return
	}

	console.log('there has been some mistake this should not be reached')

}

function saveProfileData() {

	var num
	if (gender.val() == "1") {
		num = 0
	}
	else if (gender.val() == "2") {
		num = 1
	}
	else {
		num = 0
	}
	var prof = {
		"username" : username.val(),
		"email" : email.val(),
		"name" : name.val(),
		"lastname" : lastname.val(),
		"phone" : phone.val(),
		"gender" : num,
		"dateofbirth" : dateofbirth.val(),
		"website" : website.val(),
		"biography" : bio.val()		
	}

	var dataz = JSON.stringify(prof)
	console.log(prof)
	

	$.ajax({
		type:'POST',
		crossDomain: true,
		url: PROFILE_SERVICE_URL + '/editprofile',
		contentType : 'application/json',
		dataType: 'JSON',
		data : dataz,
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success : function(data) {
		    console.log(data)
		    window.location.reload()
		},
		error : function(xhr, status, data) {
		    console.log(xhr)
		    console.log('Cant save profile data');
		}
	    })

	

}

function savePrivacyData() {

	var privacy = {
		"is_public" : profprivacy.is(':checked') ,
		"is_inbox_open" : dmprivacy.is(':checked'),
		"is_tagging_allowed" : tagprivacy.is(':checked')
	}
	console.log(privacy)

	var dataz = JSON.stringify(privacy)

	// /editprivacysettings

	$.ajax({
		type:'POST',
		crossDomain: true,
		url: PROFILE_SERVICE_URL + '/editprivacysettings',
		contentType : 'application/json',
		dataType: 'JSON',
		data : dataz,
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success : function(data) {
		    console.log(data)
		    window.location.reload()
		},
		error : function(xhr, status, data) {
		    console.log(xhr)
		    console.log('Cant save profile data');
		}
	    })
}

function saveNotificationData() {

	var notification = {
		"show_follow_notification" : follownotif.is(':checked') ,
		"show_dm_notification" : dmnotif.is(':checked'),
		"show_tagged_notification" : tagnotif.is(':checked')
	}
	console.log(notification)

	var dataz = JSON.stringify(notification)

	$.ajax({
		type:'POST',
		crossDomain: true,
		url: PROFILE_SERVICE_URL + '/editnotifsettings',
		contentType : 'application/json',
		dataType: 'JSON',
		data : dataz,
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success : function(data) {
		    console.log(data)
		    window.location.reload()
		},
		error : function(xhr, status, data) {
		    console.log(xhr)
		    console.log('Cant save profile data');
		}
	    })

}

function saveMuteBlockData() {
	console.log('not yet')
}