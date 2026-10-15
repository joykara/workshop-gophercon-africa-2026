import { mount } from '/assets/js/threejs.js';

// A slow-turning globe with glowing sighting hotspots. The points come
// from a <script type="application/json" id="globe-points"> emitted by
// templ.JSONScript, each one a [lat, lng] pair in degrees.
mount('globe', function ({ THREE, scene, camera, renderer, canvas, pointer, colour, still }) {
	const base = colour('--globe-sphere', '#2d1a1f');
	const land = colour('--globe-rim', '#5a6db5');
	const spot = colour('--globe-point', '#f0a858');

	renderer.setClearColor(0x000000, 0);

	camera.position.set(0, 0, 3.2);

	const ambient = new THREE.AmbientLight(0xffffff, 0.55);
	scene.add(ambient);
	const key = new THREE.DirectionalLight(0xffffff, 0.9);
	key.position.set(2, 1.5, 3);
	scene.add(key);

	const group = new THREE.Group();
	scene.add(group);

	// The sphere itself: matte and dark, so hotspots glow against it.
	const sphere = new THREE.Mesh(
		new THREE.SphereGeometry(1, 64, 64),
		new THREE.MeshStandardMaterial({ color: base, roughness: 0.95, metalness: 0.05 })
	);
	group.add(sphere);

	// A faint wire shell for the "cool futuristic" grid feel.
	const wire = new THREE.Mesh(
		new THREE.SphereGeometry(1.004, 32, 24),
		new THREE.MeshBasicMaterial({ color: land, wireframe: true, transparent: true, opacity: 0.18 })
	);
	group.add(wire);

	// Glow halo around the planet.
	const halo = new THREE.Mesh(
		new THREE.SphereGeometry(1.08, 48, 48),
		new THREE.MeshBasicMaterial({ color: spot, transparent: true, opacity: 0.08, side: THREE.BackSide })
	);
	group.add(halo);

	// Points for each sighting location.
	const data = JSON.parse(document.getElementById('globe-points')?.textContent || '[]');
	if (data.length) {
		const positions = new Float32Array(data.length * 3);
		for (let i = 0; i < data.length; i++) {
			const latDeg = Number(data[i][0]) || 0;
			const lngDeg = Number(data[i][1]) || 0;
			const lat = latDeg * Math.PI / 180;
			const lng = lngDeg * Math.PI / 180;
			const r = 1.015;
			positions[i * 3 + 0] = r * Math.cos(lat) * Math.sin(lng);
			positions[i * 3 + 1] = r * Math.sin(lat);
			positions[i * 3 + 2] = r * Math.cos(lat) * Math.cos(lng);
		}
		const geo = new THREE.BufferGeometry();
		geo.setAttribute('position', new THREE.BufferAttribute(positions, 3));
		const mat = new THREE.PointsMaterial({ color: spot, size: 0.028, sizeAttenuation: true, transparent: true, opacity: 0.95 });
		group.add(new THREE.Points(geo, mat));
	}

	// Start tilted so the equator is visible.
	group.rotation.x = 0.35;

	if (still) {
		return function () {};
	}

	return function (time) {
		group.rotation.y = time * 0.08;
		// Subtle parallax from the pointer, capped so it stays composed.
		group.rotation.x = 0.35 + pointer.y * 0.1;
	};
});
