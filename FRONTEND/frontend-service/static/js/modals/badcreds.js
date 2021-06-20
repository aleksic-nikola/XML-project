const login_modal_title = $('#login_modal_title')
const login_modal_body = $('#login_modal_body')
const login_alert_box = $('#login_alert_box')



function showLoginError(title, text) {
	login_modal_title.html(title)
	login_modal_body.html(text)
	login_alert_box.modal('show');
}
