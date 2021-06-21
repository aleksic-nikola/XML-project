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

	var prof = {
		username : username.val(),
		email : email.val(),
		lastname : lastname.val(),
		phone : phone.val(),
		gender : gender.val(),
		dateofbirth : dateofbirth.val(),
		website : website.val(),
		biography : bio.val()		
	    }

	console.log(prof)

}

function savePrivacyData() {

	var privacy = {
		"is_public" : profprivacy.is(':checked') ,
		"is_inbox_open" : dmprivacy.is(':checked'),
		"is_tagging_allowed" : tagprivacy.is(':checked')
	}
	console.log(privacy)
}

function saveNotificationData() {

	var notification = {
		"is_public" : follownotif.is(':checked') ,
		"is_inbox_open" : dmnotif.is(':checked'),
		"is_tagging_allowed" : tagnotif.is(':checked')
	}
	console.log(notification)
}

function saveMuteBlockData() {
	console.log('not yet')
}