'use strict';

function resizeHalf() {
    let img = document.getElementById("img1");
    img.style.width = (img.width / 2) + "px";
    img.style.height = (img.height / 2) + "px";
}

function resizeNormal() {
    let img = document.getElementById("img1");
    let width = img.naturalWidth;
    let height = img.naturalHeight;
    img.style.width = width + "px";
    img.style.height = height + "px";
}

function changePhotoJS(name, caption, width, height) {
    let img = document.getElementById("img1");
    img.src = name;
    img.style.width = width + "px";
    img.style.height = height + "px";
    document.getElementById("caption").textContent = caption;
}

function grayScaleJS() {
    let img = document.getElementById("img1");
    grayScaleGo(img.src);
}

function loadGrayScale(filename) {
    let img = document.getElementById("img1");
    img.src = filename;
}