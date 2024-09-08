document.addEventListener("DOMContentLoaded", (event) => {
    document.addEventListener("htmx:beforeSwap", (evt) => {
        if (evt.detail.xhr.status === 422) {
            evt.detail.shouldSwap = true;
            evt.detail.isError = false;
        }
    })
})