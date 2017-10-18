var User, order,JSONorders;

// function initFirebase() {
firebase.initializeApp({
    apiKey: "AIzaSyBok3MkSuQ9T-JicJ1WLPznV2em1QqO9ag",
    authDomain: "cubbychaser-8c5e4.firebaseapp.com",
    databaseURL: "https://cubbychaser-8c5e4.firebaseio.com",
    projectId: "cubbychaser-8c5e4",
    storageBucket: "",
    messagingSenderId: "900701069600"
});
    
firebase.auth().onAuthStateChanged(function (user) {
    var win = window.location.href;
    if (!user) {
        if (win.includes("main")) {
            window.location.href = "index.html";
        }
        return
    }
    if (!win.includes("main")) {
        window.location.href = "main.html";
    }
    document.getElementById('display-image').src = user.photoURL
	document.getElementById('dislplay-name').innerText = user.displayName
    User = user
});
// }

function populateSess(data) {
    var allHTML = '';
    JSONorders = data;
    for(var i in data){
        var sessHTML = '<div class="mdl-card card-square mdl-shadow--2dp session-card">' +
            '<div id="sessBar-'+i+'" class="ripplelink rip-session" >' +
            '<h5><span class="session-icon">' + data[i] + '</span> Session ' + i + '</h5>' +
            '</div></div>';
        allHTML += sessHTML;
    }
    document.getElementById("session-data").innerHTML = allHTML;
    rippleCopy();
} 

function fixRip(id) {
    document.getElementById('sessBar-'+id).classList.remove('animate');
}


function rippleCopy(){
    var links = document.querySelectorAll('.ripplelink');
    for (var i = 0, len = links.length; i < len; i++) {
        links[i].addEventListener('click', function(e) {
            var targetEl = e.target;
            var inkEl = targetEl.querySelector('.ink');
            if (inkEl) {
                inkEl.classList.remove('animate');
            } else {
                inkEl = document.createElement('span');
                inkEl.classList.add('ink');
                inkEl.style.width = inkEl.style.height = Math.max(targetEl.offsetWidth, targetEl.offsetHeight) +
                    'px';
                targetEl.appendChild(inkEl);
            }

            inkEl.style.left = (e.offsetX - inkEl.offsetWidth / 2) + 'px';
            inkEl.style.top = (e.offsetY - inkEl.offsetHeight / 2) + 'px';
            inkEl.classList.add('animate');
        }, false);
    }

    var clipboard = new Clipboard(links);

    clipboard.on('success', function(e) {
        console.log(e);
    });
    clipboard.on('error', function(e) {
        console.log(e);
    });
}




function showSesses(){
    document.querySelector("#session-dialog").showModal();
}

function pad(n, width, z) {
    z = z || '0';
    n = n + '';
    return n.length >= width ? n : new Array(width - n.length + 1).join(z) + n;
  }

function populateCubbies(full) {
    // clearCubbies();
    for (var i in full.Cubbies) {
        var loc = padSpot(i);
        document.getElementById("rip-" + loc).setAttribute("data-clipboard-text", full.Cubbies[i].OrderNumber);
    }
    document.getElementById("sess-drop").innerHTML = '<i class="material-icons">keyboard_arrow_down</i> Session '+full.ID;
    document.getElementById("end-sess").removeAttribute("disabled");
    document.getElementById("show-sess").setAttribute("disabled", '');
    closeSess();
    fixRip(full.ID);
    go2you()
}

function clearCubbies(){
    var clear = document.getElementsByClassName("clear-class");
    var noCopy = document.getElementsByClassName("ripplelink");
    var band = document.getElementsByClassName("mdl-card__actions");
    var endSession = document.getElementById("end-sess").setAttribute("disabled","");

    [].slice.call(clear).forEach(function(clear){
        clear.innerHTML = "";
    });
    [].slice.call(noCopy).forEach(function(noCopy){
        noCopy.removeAttribute("data-clipboard-text");
    });
    [].slice.call(band).forEach(function(noCopy){
        noCopy.removeAttribute("style");
    });
    document.getElementById("sess-drop").innerHTML = '<i class="material-icons">keyboard_arrow_down</i> Sessions ';
}

function logout() {
    firebase.auth().signOut().then(function () {}, function (error) {
        alert("You did not log out for some reason.");
    });
}

// function skpcSearch(e) {
//     if (e.keyCode == 13) {
//         var skupc = document.getElementById("upc-sku");
//         // Replace with function to server
//         console.log(skupc.value);
//         order = fakeSend(skupc.value, User.email);
//         // Replace above
//         if (!order) {
//             cubbyLoc = 0;
//             alertMaterial(skupc);
//             skupc.focus();
//             return false;
//         }
//         sendToCubby(spot, sess, full);
//         document.getElementById("cubby").focus();
//         return false;
//     }
// }

function preloadImages(cubs) {
    for (var s in cubs) {
        for (var i in cubs[s].Items) {
            var f = new Image();
            f.src = cubs[s].Items[i].ImageURL
            // console.log(f.src)
        }
    }
}

function sendToCubby(img, spot) {
    go2cub();
    var loc = padSpot(spot);

    document.getElementById("prog-" + loc).style.visibility = 'hidden'
    
    var cub = document.getElementById(loc);
    var cubImg = document.getElementById("img-" + loc);
    var cubBan = document.getElementById("band-" + loc);
    cubImg.innerHTML = '<img src="' + img + '" alt="" >';

    cub.classList.add("tada");
    // cub.addEventListener('click', function(e) { e.stopPropagation() })
}

function go2you() {
    var upcSKU = document.getElementById("upc-sku")
    upcSKU.value = ''
    document.getElementById("cubby").value = ''
    upcSKU.focus()
}
function clryou() {
    document.getElementById("upc-sku").value = ''
}
function clrcub() {
    document.getElementById("cubby").value = ''
}
function go2cub() {
    document.getElementById("cubby").focus()
}

function stopShake(spot, qt, tot) {
    document.getElementById(padSpot(spot)).classList.remove("tada");
    document.getElementById('clickwall').setAttribute('hidden', '');
    orderCount(spot, qt, tot)
    go2you()
}

function cubbyAssign(e) {
    // if (e.keyCode == 13) {
    //     var skupc = document.getElementById("upc-sku");
    //     var cubby = document.getElementById("cubby");
    //     var pat = /[Dd][0-9]{3,}/;
    //     if (!pat.test(cubby.value)) {
    //         return;
    //     }
    //     if (order.location != cubby.value) {
    //         alertMaterial(cubby);
    //         return;
    //     }
    //     // Replace with fucntion to send to server
    //     // also make true/false for errors to move on
    //     var ok = fakeSend(cubby.value, User.email);
    //     document.getElementById(order.location).classList.remove("tada");
    //     orderCount(spot, sess, full);
    //     console.log(cubby.value);
    //     // Replace above
    //     if (!ok) {
    //         alert("An error has occrued on the server");
    //         return;
    //     }
    //     skupc.value = '';
    //     cubby.value = '';
    //     skupc.focus();
    //     return;
    // }
}

function padSpot(spot) {
    return 'D'+pad(Number(spot)+1,4)
}

function orderCount(spot, qt, tot) {
    var cl = document.getElementById(padSpot(spot)).classList
    if (cl.contains('tada')) {
        return
    }

    var loc = padSpot(spot);
    var cubbyImg = document.getElementById("img-" + loc);
    if (qt == tot) {
        cubbyImg.innerHTML = '<i class="material-icons cubby-done">done</i>';
    } else {
        cubbyImg.innerHTML = '<span class="order-count">' + qt + "/" + tot + '</span>';
    }
    
    var cubProg = document.getElementById("prog-" + loc);
    var wid = 100*qt / tot
    if (wid == 0) {
        cubProg.style.visibility = 'hidden'
    } else {
        cubProg.style.width = wid+"%";
        cubProg.style.visibility = 'visible'
    }
}

function fakeSend(val, user) {
    if (val.includes("D")) {
        return true;
    }
    if (val == "21") {
        console.log(user);
        return {
            location: "D002",
            url: "https://s3.amazonaws.com/wedgenix-host/CW-762166276148+(1)-NEW.jpg",
            total: "3",
            spot: "1"
        };
    }
    if (val == "22") {
        console.log(user);
        return {
            location: "D002",
            url: "https://s3.amazonaws.com/wedgenix-host/829102001131+(1)-NEW.jpg",
            total: "3",
            spot: "2"
        };
    }
    if (val == "23") {
        console.log(user);
        return {
            location: "D002",
            url: "https://s3.amazonaws.com/wedgenix-host/CW-762166374158 (1).jpg",
            total: "3",
            spot: "3"
        };
    }

    if (val == "11") {
        console.log(user);
        return {
            location: "D001",
            url: "https://s3.amazonaws.com/wedgenix-host/CW-762166276148+(1)-NEW.jpg",
            total: "3",
            spot: "1"
        };
    }
    if (val == "12") {
        console.log(user);
        return {
            location: "D001",
            url: "https://s3.amazonaws.com/wedgenix-host/829102001131+(1)-NEW.jpg",
            total: "3",
            spot: "2"
        };
    }
    if (val == "13") {
        console.log(user);
        return {
            location: "D001",
            url: "https://s3.amazonaws.com/wedgenix-host/CW-762166374158 (1).jpg",
            total: "3",
            spot: "3"
        };
    }
    return false;
}

function alertMaterial(elem) {
    var val = elem.value;
    var title, msg;
    if (elem.id != "cubby") {
        title = "Wrong SKU";
        msg = "<b>" + val + "</b> is not a known SKU. Try using UPC or re-enter the SKU.";
        var UPCreg = /[0-9]{12,13}$/;
        if (val == '') {
            title = "Missing UPC/SKU";
            msg = "You must provide a UPC or SKU to continue";
        }

        if (UPCreg.test(val)) {
            title = "Wrong UPC";
            msg = "<b>" + val + "</b> is not a known UPC. Try using SKU or re-enter the UPC.";
        }
    } else {
        title = "Wrong Cubby";
        msg = "<b>" + val + "</b> is not the right cubby. Check that you put the product into the right cubby.";
    }
    var dialog = document.querySelector('#dialog');
    document.getElementById("warning-title").innerHTML = title;
    document.getElementById("warning-message").innerHTML = msg;
    dialog.showModal();
}

function materialAlert(title, msg) {
    var dialog = document.querySelector('#dialog');
    document.getElementById("warning-title").innerHTML = title;
    document.getElementById("warning-message").innerHTML = msg;
    dialog.showModal();
}

function login() {
    var provider = new firebase.auth.GoogleAuthProvider();
    provider.addScope('https://www.googleapis.com/auth/contacts.readonly');

    firebase.auth().signInWithPopup(provider).then(function (result) {
        // This gives you a Google Access Token. You can use it to access the Google API.
        var token = result.credential.accessToken;
        // The signed-in user info.
        var user = result.user;
        // ...
        console.log(user);
    }).catch(function (error) {
        // Handle Errors here.
        var errorCode = error.code;
        var errorMessage = error.message;
        // The email of the user's account used.
        var email = error.email;
        // The firebase.auth.AuthCredential type that was used.
        var credential = error.credential;
        // ...
        console.log(errorCode);
    });
}

function closeDialog() {
    var title = document.getElementById("warning-title").innerHTML;
    var skupc = document.getElementById("upc-sku");
    var cubby = document.getElementById("cubby");
    dialog.close();
    if (title.includes("Cubby")) {
        cubby.focus();
    } else {
        skupc.focus();
    }
}
function closeSess() {
    document.querySelector("#session-dialog").close();
}




// FOR TESTING ONLY
function initTestData() {

    return JSON.parse('{' +
        '"22":{' +
        '"D001": {' +
        '"items": {' +
        '"11": {' +
        '"url:": "https://s3.amazonaws.com/wedgenix-host/CW-762166276148+(1)-NEW.jpg",' +
        '"spot": "1"' +
        '},' +
        '"12": {' +
        '"url": "https://s3.amazonaws.com/wedgenix-host/829102001131+(1)-NEW.jpg",' +
        '"spot": "2"' +
        '},' +
        '"13": {' +
        '"url": "https://s3.amazonaws.com/wedgenix-host/CW-762166374158 (1).jpg",' +
        '"spot": "3"' +
        '}' +
        '},' +
        '"orderNumber": "111-222-333",' +
        '"total": "3"' +
        '},' +
        '"D002": {' +
        '"items": {' +
        '"21": {' +
        '"url:": "https://s3.amazonaws.com/wedgenix-host/CW-762166276148+(1)-NEW.jpg",' +
        '"spot": "1"' +
        '},' +
        '"22": {' +
        '"url": "https://s3.amazonaws.com/wedgenix-host/829102001131+(1)-NEW.jpg",' +
        '"spot": "2"' +
        '},' +
        '"23": {' +
        '"url": "https://s3.amazonaws.com/wedgenix-host/CW-762166374158 (1).jpg",' +
        '"spot": "3"' +
        '}' +
        '},' +
        '"total": "3",' +
        '"orderNumber": "222-333-999"' +
        '},' +
        '"count":"6"'+
        '},' +
        '"33":{' +
        '"D001": {' +
        '"items": {' +
        '"11": {' +
        '"url:": "https://s3.amazonaws.com/wedgenix-host/CW-762166276148+(1)-NEW.jpg",' +
        '"spot": "1"' +
        '},' +
        '"12": {' +
        '"url": "https://s3.amazonaws.com/wedgenix-host/829102001131+(1)-NEW.jpg",' +
        '"spot": "2"' +
        '},' +
        '"13": {' +
        '"url": "https://s3.amazonaws.com/wedgenix-host/CW-762166374158 (1).jpg",' +
        '"spot": "3"' +
        '}' +
        '},' +
        '"orderNumber": "555-555-555",' +
        '"total": "3"' +
        '},' +
        '"D002": {' +
        '"items": {' +
        '"21": {' +
        '"url:": "https://s3.amazonaws.com/wedgenix-host/CW-762166276148+(1)-NEW.jpg",' +
        '"spot": "1"' +
        '},' +
        '"22": {' +
        '"url": "https://s3.amazonaws.com/wedgenix-host/829102001131+(1)-NEW.jpg",' +
        '"spot": "2"' +
        '},' +
        '"23": {' +
        '"url": "https://s3.amazonaws.com/wedgenix-host/CW-762166374158 (1).jpg",' +
        '"spot": "3"' +
        '}' +
        '},' +
        '"total": "3",' +
        '"orderNumber": "888-668-999"' +
        '},' +
        '"D003": {' +
        '"items": {' +
        '"31": {' +
        '"url:": "https://s3.amazonaws.com/wedgenix-host/CW-762166276148+(1)-NEW.jpg",' +
        '"spot": "1"' +
        '},' +
        '"32": {' +
        '"url": "https://s3.amazonaws.com/wedgenix-host/829102001131+(1)-NEW.jpg",' +
        '"spot": "2"' +
        '},' +
        '"33": {' +
        '"url": "https://s3.amazonaws.com/wedgenix-host/CW-762166374158 (1).jpg",' +
        '"spot": "3"' +
        '}' +
        '},' +
        '"total": "3",' +
        '"orderNumber": "999-999-999"' +
        '},' +
        '"count":"9"'+
        '}' +
        '}');


}