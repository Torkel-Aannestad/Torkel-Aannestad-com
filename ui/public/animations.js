function animateOnLoad() {
  gsap.from(".gsap-fade-in", {
    duration: 1,
    y: 150,
    opacity: 0,
    ease: "expo.out",
  });

  //on load
  gsap.from(".gsap-ohoi", {
    delay: 1,
    duration: 1.5,
    y: 150,
    scale: 0,
    opacity: 0,
    ease: "expo.out",
  });
}
document.addEventListener("DOMContentLoaded", animateOnLoad);

const animateBubble = gsap.to(".gsap-ohoi", {
  paused: true,
  scale: 1.1,
  duration: 0.1,
});
const bubble = document.querySelector(".gsap-h1");
bubble.addEventListener("mouseenter", () => animateBubble.play());
bubble.addEventListener("mouseleave", () => animateBubble.reverse());
