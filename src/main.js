'use strict';
console.log('Starting');

console.log('Loading main.js');
require('./screeps');

const wasmTick = true

module.exports = {
    loop: () => {
        const startCPU = Game.cpu.getUsed()
        const stats = Game.cpu.getHeapStatistics()
        if (wasmTick)
            global.go.loop()
        console.log("<i style='opacity: 0.5'><small>WASM tick took " + (Game.cpu.getUsed() - startCPU).toFixed(2) + "cpu</small></i>")
        console.log("<i style='opacity: 0.5'><small>Heap memory " + (stats.used_heap_size / stats.total_heap_size).toFixed(2) + "k</small></i>")
        console.log("<i style='opacity: 0.5'><small>----------- " + (stats.total_heap_size / stats.heap_size_limit).toFixed(2) + "k</small></i>")
    }
}