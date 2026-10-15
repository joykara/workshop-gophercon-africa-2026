// Boots the hero carousel: full-width slides, autoplay left-to-right,
// prev/next arrows, pagination bullets. Pauses on hover and when the
// reader has asked for reduced motion. Also handles the 3D hover tilt
// on the active slide's photo.

(function () {
	var el = document.querySelector('.hero__swiper');
	if (!el || typeof Swiper === 'undefined') return;

	var reduceMotion = window.matchMedia('(prefers-reduced-motion: reduce)').matches;

	new Swiper(el, {
		slidesPerView: 1,
		spaceBetween: 0,
		loop: true,
		speed: 900,
		autoplay: reduceMotion ? false : { delay: 5500, pauseOnMouseEnter: true, disableOnInteraction: false },
		keyboard: { enabled: true },
		pagination: { el: el.querySelector('.hero__pagination'), clickable: true },
		navigation: {
			nextEl: el.querySelector('.hero__nav--next'),
			prevEl: el.querySelector('.hero__nav--prev'),
		},
	});

	if (reduceMotion) return;

	// 3D hover tilt — tracks the pointer across the hero and tilts the
	// active slide's photo-wrap. The transform itself is in CSS; this
	// only sets --tilt-x / --tilt-y custom properties.
	var hero = el.closest('.hero');
	if (!hero) return;
	var maxTilt = 4; // degrees — subtle, not party-trick

	hero.addEventListener('pointermove', function (e) {
		var rect = hero.getBoundingClientRect();
		var dx = (e.clientX - rect.left) / rect.width - 0.5;
		var dy = (e.clientY - rect.top) / rect.height - 0.5;
		var wrap = hero.querySelector('.swiper-slide-active .hero__photo-wrap');
		if (!wrap) return;
		wrap.style.setProperty('--tilt-x', (-dy * maxTilt).toFixed(2) + 'deg');
		wrap.style.setProperty('--tilt-y', (dx * maxTilt).toFixed(2) + 'deg');
	});

	hero.addEventListener('pointerleave', function () {
		hero.querySelectorAll('.hero__photo-wrap').forEach(function (w) {
			w.style.setProperty('--tilt-x', '0deg');
			w.style.setProperty('--tilt-y', '0deg');
		});
	});
})();

// Same 3D tilt on the sighting-detail hero.
(function () {
	if (window.matchMedia('(prefers-reduced-motion: reduce)').matches) return;
	var el = document.querySelector('.detail-hero[data-tilt]');
	if (!el) return;
	var wrap = el.querySelector('.detail-hero__photo-wrap');
	if (!wrap) return;
	var maxTilt = 3;
	el.addEventListener('pointermove', function (e) {
		var rect = el.getBoundingClientRect();
		var dx = (e.clientX - rect.left) / rect.width - 0.5;
		var dy = (e.clientY - rect.top) / rect.height - 0.5;
		wrap.style.setProperty('--tilt-x', (-dy * maxTilt).toFixed(2) + 'deg');
		wrap.style.setProperty('--tilt-y', (dx * maxTilt).toFixed(2) + 'deg');
	});
	el.addEventListener('pointerleave', function () {
		wrap.style.setProperty('--tilt-x', '0deg');
		wrap.style.setProperty('--tilt-y', '0deg');
	});
})();
