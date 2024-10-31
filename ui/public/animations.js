function animate() {
  gsap.from(".gsap-fade-in", {
    duration: 1,
    y: 150,
    opacity: 0,
    ease: "expo.out",
  });
}

document.addEventListener("DOMContentLoaded", animate);
