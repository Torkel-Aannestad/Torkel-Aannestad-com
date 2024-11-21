function animateOnLoad() {
  gsap.from(".gsap-fade-in", {
    duration: 1,
    y: 100,
    opacity: 0,
    ease: "expo.out",
  });

  const ohoiElements = document.querySelectorAll(".gsap-ohoi");
  if (ohoiElements.length > 0) {
    gsap.from(".gsap-ohoi", {
      delay: 1,
      duration: 1.5,
      y: 150,
      scale: 0,
      opacity: 0,
      ease: "expo.out",
    });
  }
}
document.addEventListener("DOMContentLoaded", animateOnLoad);
document.addEventListener("htmx:afterSwap", animateOnLoad);
