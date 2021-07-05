const search_field = $('#search_input')
var public_profiles = []
var allowed_profiles = []
var public_names = {}
var allowed_names = {}
var final_profiles = []
var final_posts = []
var current_input = 'valjda_nece_ovo_nikad_upisati13151q6161q67'
const profile_list = $('#profile_list')
var location_counter = 0
var tag_counter = 0



$(document).ready(function() {

	who_am_I()
}) 

/*
search_field.focusin(function() {

	console.log('focus')
	if (search_field.val() == current_input)
		return 
	current_input = search_field.val()
	who_am_I()
})
*/

search_field.keyup(function(){

	if (search_field.val().includes('^'))
    	filterByLocation() 
	else if (search_field.val().includes('#'))
		filterByTags()
	else if (search_field.val().includes('@'))
		filterUsers()	
	else  {
		console.log('filtering all categories')
		filterByEverything()
	}

})

function filterUsers() {

	const searchInput = search_field.val().toLowerCase().trim();
	const searchString = searchInput.replaceAll('@', '').trim()
	//console.log(searchString)
	const filtered_users = final_profiles.filter((profile) => {
		return (
			profile.username.toLowerCase().includes(searchString)
			);
		});
	fillColumnsProfiles(filtered_users);

}

function filterByEverything() {
	emptyByTags()
	emptyByLocation()
	const searchString = search_field.val().toLowerCase().trim();
	//console.log(searchString)
	const filtered_posts = final_posts.filter((post) => {
		return (
			post.description.toLowerCase().includes(searchString) || 
			post.medias[0].location.city.toLowerCase().includes(searchString) ||
			post.medias[0].location.country.toLowerCase().includes(searchString) ||
			post.postedby.toLowerCase().includes(searchString)
			);
		});
	
	filterUsers()
	fillColumnsTags(filtered_posts, false);
	fillColumnsLocation(filtered_posts, false)



}

function filterByLocation() {
	emptyByLocation()
	const searchInput = search_field.val().toLowerCase().trim();
	const searchString = searchInput.replaceAll('^', '').trim()
	//console.log(searchString)
	const filtered_posts = final_posts.filter((post) => {
		return (
			post.medias[0].location.city.toLowerCase().includes(searchString) ||
			post.medias[0].location.country.toLowerCase().includes(searchString) 
			);
		});
	fillColumnsLocation(filtered_posts, false);

}

function filterByTags() {
	emptyByTags()
	const searchString = search_field.val().toLowerCase().trim();
	//console.log(searchString)
	const filtered_posts = final_posts.filter((post) => {
		return (
			post.description.toLowerCase().includes(searchString)
			);
		});
	fillColumnsTags(filtered_posts, false);


}


function who_am_I() {

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
		    console.log('using authenticated search..')
		    authenticatedSearch(data)
		},
		error: function () {
		    console.log('using non authenticated search..')
		    nonAuthenticatedSearch()
		}
	})

}

function authenticatedSearch() {
	getAllPublicProfiles()
	getAllPublicPosts()
	getAllAllowedProfiles()
	getAllAllowedPosts()
	createFinalProfileList()
}

function createFinalProfileList() {
	
}

function nonAuthenticatedSearch() {
	getAllPublicProfiles()
	//getAllPublicPosts()
}

function getAllPublicProfiles() {

	$.ajax({
		type: 'GET',
		crossDomain: true,
		url: PROFILE_SERVICE_URL + '/publicprofiles',
		contentType: 'application/json',
		dataType: 'JSON',
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success: function (data) {
		    //console.log(data)
			data.forEach(function(d) {
				if (final_profiles.filter(e => e.username == d.username).length == 0) {
					final_profiles.push(d)
				}		  
			})
			console.log('Size of final_profiles is ' + final_posts.length)
			//public_profiles.push.apply(public_profiles, data)
			fillColumnsProfiles(final_profiles)
		    getAllPublicPosts(data)
		},
		error: function () {
		    alert('error')
		}
	})

}

function getAllPublicPosts(profiles) {
	console.log('profiles is')
	console.log(profiles)
	profiles.forEach(function(p) {

		$.ajax({
			type: 'GET',
			crossDomain: true,
			url: CONTENT_SERVICE_URL + '/getpostsbyuser/' + p.username,
			contentType: 'application/json',
			dataType: 'JSON',
			beforeSend: function (xhr) {
			    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
			},
			success: function (data) {
			    console.log(data)
			    fillColumnsLocation(data, true)
				fillColumnsTags(data, false)
			},
			error: function () {
			    alert('error')
			}
		})

	})

}

function getAllAllowedProfiles() {

}

function getAllAllowedPosts() {

}

function emptyByLocation() {
	$('#column-right-0').html('')
	$('#column-right-1').html('')
	$('#column-right-2').html('')
	$('#column-right-3').html('')
	location_counter = 0
}

function emptyByTags() {
	$('#column-left-0').html('')
	$('#column-left-1').html('')
	$('#column-left-2').html('')
	$('#column-left-3').html('')
	tag_counter = 0
}

function fillColumnsLocation(pics, sentByAjax) {
	console.log('%c showing by location', 'color:blue')
	console.log(pics)
	if (sentByAjax)
		final_posts.push.apply(final_posts, pics)

	var column 
	var html = ''
	var coldiv
	for(let i = 0; i<pics.length; i++) {
		
		column = "#column-right-" + location_counter % 4
		location_counter = location_counter + 1
		coldiv = $(column)
		coldiv.append(`<img src="${pics[i].medias[0].path}">`)
		//coldiv.append(testList[i])
	}

}

function fillColumnsTags(pics, sentByAjax) {
	console.log('%c showing by location', 'color:blue')
	console.log(pics)
	if (sentByAjax)
		final_posts.push.apply(final_posts, pics)
	
	var column 
	var html = ''
	var coldiv
	for(let i = 0; i<pics.length; i++) {
		
		column = "#column-left-" + tag_counter % 4
		tag_counter = tag_counter + 1
		coldiv = $(column)
		coldiv.append(`<img src="${pics[i].medias[0].path}">`)
		//coldiv.append(testList[i])
	}

}

function fillColumnsProfiles(profiles) {
	profile_list.html('')
	// <li class="list-group-item">Cras justo odio</li>
	profiles.forEach(function(p){

		profile_list.append(`<li class="list-group-item" id="user-${p.username}" onclick="redirectToUser(this.id)">${p.username}</li>`)
	})
}

function redirectToUser(id) {
	console.log(id)
}