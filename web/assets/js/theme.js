// Toggles between the light and dark themes, remembering the choice.
(function () {
	var btn = document.getElementById('theme-toggle');
	if (!btn) return;

	function current() {
		var set = document.documentElement.getAttribute('data-theme');
		if (set) return set;
		return window.matchMedia('(prefers-color-scheme: light)').matches ? 'light' : 'dark';
	}

	function label(mode) {
		var span = btn.querySelector('.theme-toggle__label');
		if (span) span.textContent = mode === 'light' ? 'Dark' : 'Light';
		btn.setAttribute('aria-label', mode === 'light' ? 'Switch to dark theme' : 'Switch to light theme');
	}

	label(current());

	btn.addEventListener('click', function () {
		var next = current() === 'light' ? 'dark' : 'light';
		document.documentElement.setAttribute('data-theme', next);
		try { localStorage.setItem('theme', next); } catch (e) {}
		label(next);
	});
})();
