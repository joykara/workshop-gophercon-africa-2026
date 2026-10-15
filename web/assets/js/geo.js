// Switches the "Globe / Map" tabs on the home geography section.
(function () {
	document.querySelectorAll('[data-geo]').forEach(function (geo) {
		var tabs = geo.querySelectorAll('.geo__tab');
		var views = geo.querySelectorAll('.geo__view');
		tabs.forEach(function (tab) {
			tab.addEventListener('click', function () {
				var target = tab.dataset.view;
				tabs.forEach(function (t) { t.setAttribute('aria-pressed', t === tab ? 'true' : 'false'); });
				views.forEach(function (v) {
					if (v.dataset.view === target) { v.setAttribute('data-active', ''); }
					else { v.removeAttribute('data-active'); }
				});
				// Leaflet needs a resize kick when its container has been hidden.
				if (target === 'map' && window.L) {
					var el = geo.querySelector('.map');
					if (el && el._leafletInstance) {
						setTimeout(function () { el._leafletInstance.invalidateSize(); }, 50);
					}
				}
			});
		});
	});
})();
