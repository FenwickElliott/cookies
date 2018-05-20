// Softcode partner specific variables:
let partnerID = "inception";
// let tracker = "http://localhost:4000";
let tracker = "https://cookies.fenwickelliott.io"
// let tracker = "https://glam-cookie-service.appspot.com"



// Workflow:
let cookieID = getCookie(partnerID + "ID");

if (!cookieID) {
    setCookie();
}

document.getElementById("trackingFrame").src = tracker + "/in?partner=" + partnerID + "&cookie=" + cookieID;

//Utility functions:

function getCookie(name) {
    let value = "; " + document.cookie;
    let parts = value.split("; " + name + "=");
    if (parts.length == 2) return parts.pop().split(";").shift();
}

function setCookie() {
    let date = new Date();
    cookieID = sha1(window + date);
    date.setTime(date.getTime()+(365*24*60*60*1000));
    document.cookie = partnerID + "ID=" + cookieID + ";expires="+ date.toUTCString();
}