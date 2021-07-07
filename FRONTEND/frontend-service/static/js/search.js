const search_field = $('#search_input')
var public_profiles = []
var allowed_profiles = []
var public_names = {}
var allowed_names = {}
var final_profiles = []
var final_profiles_with_private = [] // profiles + private followless profiles
var final_posts = []
var current_input = 'valjda_nece_ovo_nikad_upisati13151q6161q67'
const profile_list = $('#profile_list')
var location_counter = 0
var tag_counter = 0
var authenticated = false
var currentUserSearch


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
	const filtered_users = final_profiles_with_private.filter((profile) => {
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
		    authenticated = true
            //authenticatedSearch()
			getCurrentUserInformationSearch()
			$("#myProfileHref").attr('href', "profile.html?" + data.username);
			showActualNavs()
            
		},
		error: function () {
		    console.log('using non authenticated search..')
		    nonAuthenticatedSearch()
			showActualNavs()
		}
	})

}

function authenticatedSearch() {
    console.log('authenticated search in progress..')
	//getAllPublicProfiles()
	getAllAllowedProfiles()
	//getAllAllowedPosts()
	//createFinalProfileList()
}

function getAllAllowedProfiles() {

    console.log('getAllAllowedProfiles')

    $.ajax({
		type: 'GET',
		crossDomain: true,
		url: PROFILE_SERVICE_URL + '/allowedprofiles',
		contentType: 'application/json',
		dataType: 'JSON',
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success: function (data) {
		    //console.log(data)
			
			
			if (data != null) {
                console.log('getAllAllowedProfiles size is > 0')
                data.forEach(function(d) {
                    if (final_profiles.filter(e => e.username == d.username).length == 0) {
                        final_profiles.push(d)
                    }		  
                })
            } 

		
			//console.log('Size of final_profiles is ' + final_posts.length)
			//public_profiles.push.apply(public_profiles, data)
			getAllPublicProfiles()
		},
		error: function () {
		    alert('error')
		}
	})

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
			//console.log('Size of final_profiles is ' + final_posts.length)
			//public_profiles.push.apply(public_profiles, data)
			if(authenticated) {
				filterBlockedAndMutedUsers(final_profiles, false)
			}

            getAllPublicPosts(final_profiles)
            final_profiles_with_private.push.apply(final_profiles_with_private, final_profiles)

            if (authenticated) {
                getPrivateNonFollowed()
            }
            else {
                fillColumnsProfiles(final_profiles)
            }
		},
		error: function () {
		    alert('error')
		}
	})

}

// get all private profiles that the current user does nto follow
function getPrivateNonFollowed() {

    $.ajax({
		type: 'GET',
		crossDomain: true,
		url: PROFILE_SERVICE_URL + '/followlessprivateprofiles',
		contentType: 'application/json',
		dataType: 'JSON',
		beforeSend: function (xhr) {
		    xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
		},
		success: function (data) {
		    console.log('using authenticated search..')
            data.forEach(function(d) {
				if (final_profiles_with_private.filter(e => e.username == d.username).length == 0) {
					final_profiles_with_private.push(d)
				}		  
			})
            //fillColumnsProfiles(final_profiles_with_private)
			if(authenticated) {
				filterBlockedAndMutedUsers(final_profiles_with_private, true)
			}

			searchquery()
			
		},
		error: function () {
		    console.log('using non authenticated search..')
		    nonAuthenticatedSearch()
		}
	})

}

// gets all public posts (only public for non logged in and public + followed for logged in)
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

				if(!authenticated) {
					searchquery()
				}

			},
			error: function () {
			    alert('error')
			}
		})

	})

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

		if (pics[i].medias[0].path.split(".")[1] != 'jpg' && pics[i].medias[0].path.split(".")[1] != 'png') {
			//coldiv.append(`<img src="${pics[i].medias[0].path}" id="post-${pics[i].ID}">`)

			coldiv.append(`
			
			<video width="100%" height="100%" id="post-${pics[i].ID}" class="card-img-top gallery-item">
			<source src="${pics[i].medias[0].path}"  type="video/mp4">                       
			Your browser does not support the video tag.
			</video>
			
			
			`)
		} else {
			coldiv.append(`<img src="${pics[i].medias[0].path}" id="post-${pics[i].ID}">`)
		}


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
		//coldiv.append(`<img src="${pics[i].medias[0].path}" id="post-${pics[i].ID}">`)
		//coldiv.append(testList[i])

		if (pics[i].medias[0].path.split(".")[1] != 'jpg' && pics[i].medias[0].path.split(".")[1] != 'png') {
			//coldiv.append(`<img src="${pics[i].medias[0].path}" id="post-${pics[i].ID}">`)

			coldiv.append(`
			
			<video width="100%" height="100%" id="post-${pics[i].ID}" class="card-img-top gallery-item">
			<source src="${pics[i].medias[0].path}"  type="video/mp4">                       
			Your browser does not support the video tag.
			</video>
			
			
			`)
		} else {
			coldiv.append(`<img src="${pics[i].medias[0].path}" id="post-${pics[i].ID}">`)
		}

		console.log()
	}

}

function fillColumnsProfiles(profiles) {

	profile_list.html('')
	// <li class="list-group-item">Cras justo odio</li>
	profiles.forEach(function(p){

		profile_list.append(`<li class="list-group-item" id="user-${p.username}" onclick="redirectToUserProf(this.id)"><a href="#">${p.username}</a></li>`)
	})
}

function redirectToUser(id) {
	console.log(id)
}

document.addEventListener("click", function (e) {

        const post_id = e.target.getAttribute("id")
        console.log(post_id)

		if(post_id != null && post_id.includes("post-")) {
			showImageModalSearch(post_id.split("-")[1], final_posts)
			console.log("---------------------------")
			console.log(final_posts)

			currentOpenedPost = post_id
		} else {
			return
		}

})

const which_image = $('#which_image1')
const like_button_div = $('#like_button_div1')
const comments_body = $('#comments_body1')
const post_comment_button_div = $('#post_comment_button_div1')
const image_description = $('#image_description1')

function showImageModalSearch(id, postList) {

	// if user is logged in show images - if not he cant open image
	if(authenticated) {
		showModalForLoggedUser(id, postList)
	} else {
		showModalForNonLoggedUser(id, postList)
	}

}

function showModalForLoggedUser(id, postList) {
	console.log('showing image ' + id)
	var post = undefined
	postList.forEach(function(p) {
		console.log('post loop ' + p.ID)
		console.log("ID: " + id + "----" + "p.ID: " + p.ID)
		console.log(id==p.ID)
		console.log(id)
		console.log(p.ID)
		if (p.ID == id) {
			post = p
			console.log('post is ' + post)
		} else {
			return
		}
	})
	var media = post.medias[0].path

	//console.log("==============> length of post (num of content): " + post.medias.length)

	if (post.medias[0].path.split(".")[1] != 'jpg' && post.medias[0].path.split(".")[1] != 'png' && post.medias.length == 1) {
		// only one VIDEO

		which_image.html(`

		<video width="100%" height="100%" controls alt="Card image cap" class="card-img-top modal-img">
			<source src="${media}" type="video/mp4">                       
			Your browser does not support the video tag.
		</video>
		
		`)

	} else {
		
		if(post.medias.length > 1) {
			// MORE VIDEOS or PICTURES (albums)

			console.log("There is more then 1 PICTURE or VIDEO")

			// carousel
			var postsHTML = ''


				postsHTML +=  `<div id="demo-${post.ID}" class="carousel slide" data-ride="carousel">`
				postsHTML +=  `<ul class="carousel-indicators">`
		
				for(var j=0; j < post.medias.length ; j++) {

					if(j==0) {
						postsHTML += `<li data-target="#demo-${post.ID}" data-slide-to="${j}" class="active"></li>`
					} else {
						postsHTML += `<li data-target="#demo-${post.ID}" data-slide-to="${j}" ></li>`    
					}
		
				}
				
		
				postsHTML +=  `</ul>`
		
				// SLIDESHOW:
		
				postsHTML += `<div class="carousel-inner" style=" width:100%; height:auto !important; padding-left:13%; padding-right:13%">`
				
				for(var g=0; g < post.medias.length; g++) {
					img_url = '../' + post.medias[g].path
					media_string = `<img width="100%" height="auto; !important" class="card-img-top modal-img"  src="${img_url}" alt="Card image cap">`
			
					if (post.medias[g].path.split(".")[1] != 'jpg' && post.medias[g].path.split(".")[1] != 'png') {
						// then its a video
						media_string = `
									<video width="100px" height="100%" controls class="card-img-top modal-img" >
										<source src="${img_url}" type="video/mp4">                       
										Your browser does not support the video tag.
									</video>`
					}
		
		
					if(g==0) {
						postsHTML += `<div class="carousel-item active">`
						postsHTML +=  media_string
						postsHTML +=  `</div>`
					} else {
						postsHTML += `<div class="carousel-item">`
						postsHTML +=  media_string
						postsHTML +=  `</div>`
					}
		
				}
		
				postsHTML += `</div>`
		
				postsHTML += `
									<a class="carousel-control-prev" href="#demo-${post.ID}" data-slide="prev">
										<span class="carousel-control-prev-icon"  style="background-color: black; border: 1px white;"></span>
									</a>
									<a class="carousel-control-next" href="#demo-${post.ID}" data-slide="next" >
										<span class="carousel-control-next-icon" style="background-color: black; border: 1px white;"></span>
									</a>`
		
				postsHTML += `</div>`

			

			which_image.html(postsHTML)

		} else {
			// Only 1 picture

			which_image.html(`
	
			<img 
				class="card-img-top modal-img" 
				src="${media}" 
				alt="Card image cap">	
	
			`)
		}


	}


	console.log('give me image ' + post.medias[0])
	// like button id looks like  like-ID e.g. like-1 current user likes current post

	like_button_div.html(`<a href="" id="like-${post.ID}> Like &nbsp&nbsp&nbsp&nbsp</a>`)
	
	var comment_html = ''

	post.comments.forEach(function(c) {

		comment_html += `
			<tr>
			<td class = "first_td">
			<a class ="commentedby" href=""> ${c.postedby}: </a> 
			</td>
			<td>
			${c.text}
			</td>
			</tr>
		`
	})

	comments_body.html(comment_html)

	post_comment_button_div.html(`
	<button class="btn btn-primary postcommentbtn"
	 id="post_comment_btn-${post.ID}"
	  onclick="postComment(this.id)"> Post </button>
	`)

	image_description.html(post.description)

	$('#gallery-modal-search').modal('show');
}

function showModalForNonLoggedUser(id, postList) {

	$('#nonLoggedModal').modal('show');
}

function redirectToUserProf(idprof) {
	if(idprof != null && idprof.includes("user-")) {
		profid = idprof.split("-")[1]
	} else {
		return
	}

	window.location.href = 'profile.html?' + profid;

}

// FILTRIRANJE BLOKIRANIH/MUTIRANIH USERA


function filterBlockedAndMutedUsers(mylist, filterForUsers) {
	
	console.log(final_profiles_with_private)

	var filteredlist = [];
	var filteredblockedlist = [];
	//var allowedlistgray = [];
	var allowedlistblack = [];
	
	mylist.forEach(function(a) {

		if(a.blacklist == undefined) {
			allowedlistblack.push(a)
			return
		}

		if(a.blacklist.filter(e => e.username == currentUserSearch.username).length == 0) {
			allowedlistblack.push(a)
		}
	})

	// isfiltriraj blacklistovane korisnike
	allowedlistblack.forEach(function(d) {
		if(currentUserSearch.blacklist.filter(e => e.username == d.username).length == 0) {
			filteredblockedlist.push(d)
		}		
	})

	// isfiltriraj u graylistovane
	filteredblockedlist.forEach(function(b) {
		if(currentUserSearch.graylist.filter(e => e.username == b.username).length == 0) {
			filteredlist.push(b)
		}	
	})

	if(filterForUsers) {

		final_profiles_with_private = filteredlist

		fillColumnsProfiles(filteredlist)

		return
	}

	final_profiles = filteredlist

}

function checkIfUserBlockedMeAndRestrictSearch() {

    $.ajax({
        type: 'GET',
        crossDomain: true,
        url: PROFILE_SERVICE_URL + '/getuserwhoblockedme',
        contentType: 'application/json',
        dataType: 'JSON',
        beforeSend: function (xhr) {
            xhr.setRequestHeader('Authorization', 'Bearer ' + localStorage.getItem('myToken'));
        },
        success: function (data) {
            //console.log(data)
            checkIfImInHisBlacklistSearch(data)
        },
        error: function () {
            alert("Error in getUsersWhoBlockedMeAndRestricTheirProfile")
        }
    })

}

var listwhoblockedMe = []
function checkIfImInHisBlacklistSearch(data) {
    data.listwhoblockedme.forEach(function(p) {
		listwhoblockedMe.append(p.username)
    })
}

function getCurrentUserInformationSearch() {
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

            currentUserSearch = data
            //loadFeedContent();
            //loadStories();
			authenticatedSearch();

        },
        error : function(xhr, status, data) {
             alert('Error in getCurrentUserInformationSearch')
        }
    })
}

function searchquery() {

	var query = localStorage.getItem("query")

	console.log(query)
	if(query != "undefined") {
		search_field.val(query)
		localStorage.setItem("query", undefined)
		search_field.keyup()
	}
}

function showActualNavs() {
	if(authenticated) {
		$("#loginnav").hide()
	} else {
		$("#feednav").hide()
		$("#reqsnav").hide()
		$("#msgnav").hide()
		$("#profnav").hide()
		$("#logoutnav").hide()
	}

}