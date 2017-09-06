const DownloadButton = document.querySelector("button.download-button");
const ModalDownloadCloseButton = document.querySelector("#modal-close-button")
// the backdraft and the modal
const backdraft = document.querySelector(".back-draft")
const modal = document.querySelector(".modal")

DownloadButton.addEventListener("click", (e) => {
    e.preventDefault();
    // show the backdraft
    backdraft.classList.remove("hidden")
    backdraft.classList.add("animated", "fadeIn")
    // show the modal
    setTimeout(function () {
        modal.classList.remove("hidden")
        modal.classList.add("animated", "fadeInDown")
    }, 250)
})

ModalDownloadCloseButton.addEventListener("click", (e) => {
    e.preventDefault()
    // the backdraft and the modal
    modal.classList.remove("fadeInDown")
    modal.classList.add("fadeOutUp")
    setTimeout(function () {
        backdraft.classList.remove("fadeIn")
        backdraft.classList.add("fadeOut")
    }, 250)
    setTimeout(function () {
        backdraft.classList.remove("animated", "fadeOut")
        modal.classList.remove("animated", "fadeOutUp")
        modal.classList.add("hidden")
        backdraft.classList.add("hidden")
    }, 500)
    setTimeout(function () {
        modal.classList.add("hidden")
        backdraft.classList.add("hidden")
    }, 750)
})