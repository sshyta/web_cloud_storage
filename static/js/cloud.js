import * as THREE from 'https://cdn.jsdelivr.net/npm/three@0.121.0/build/three.module.js';
import { GLTFLoader } from 'https://cdn.jsdelivr.net/npm/three@0.121.0/examples/jsm/loaders/GLTFLoader.js';

const scene = new THREE.Scene();
const camera = new THREE.PerspectiveCamera(45, window.innerWidth / window.innerHeight, 0.1, 1000);
camera.position.z = 10;

const renderer = new THREE.WebGLRenderer({ alpha: true, antialias: true });
renderer.setClearColor(0x222222);
renderer.setSize(window.innerWidth, window.innerHeight);
document.body.appendChild(renderer.domElement);

const aLight = new THREE.AmbientLight(0x404040, 2);
scene.add(aLight);

const pLight = new THREE.PointLight(0xFFFFFF, 2);
pLight.position.set(0, 30, 7);
scene.add(pLight);

const helper = new THREE.PointLightHelper(pLight);
scene.add(helper);

let loader = new GLTFLoader();
let obj = null;

loader.load('/static/scens/scene.gltf', function (gltf) {
    obj = gltf;
    obj.scene.scale.set(1.3, 1.3, 1.3);
    obj.scene.position.set(0.02, 2, 0);
    scene.add(obj.scene);
}, function (xhr) {
    console.log((xhr.loaded / xhr.total * 100) + '% loaded');
}, function (error) {
    console.error('Error loading model:', error);
});

function animate() {
    requestAnimationFrame(animate);
    if (obj && obj.scene) {
        obj.scene.rotation.y += 0.03;
    }
    renderer.render(scene, camera);
}

animate();