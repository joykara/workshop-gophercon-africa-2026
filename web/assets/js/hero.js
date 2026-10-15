// Boots the hero carousel: full-width slides, autoplay left-to-right,
// prev/next arrows, pagination bullets. Pauses on hover and when the
// reader has asked for reduced motion.

(function () {
	var el = document.querySelector('.hero__swiper');
	if (!el || typeof Swiper === 'undefined') return;

	var reduceMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;

	new Swiper(el, {
		slidesPerView: 1,
		spaceBetween: 0,
		loop: true,
		speed: 900,
		autoplay: reduceMotion ? false : { delay: 5000, pauseOnMouseEnter: true, disableOnInteraction: false },
		keyboard: { enabled: true },
		pagination: { el: el.querySelector('.hero__pagination'), clickable: true },
		navigation: {
			nextEl: el.querySelector('.hero__nav--next'),
			prevEl: el.querySelector('.hero__nav--prev'),
		},
	});
})();
