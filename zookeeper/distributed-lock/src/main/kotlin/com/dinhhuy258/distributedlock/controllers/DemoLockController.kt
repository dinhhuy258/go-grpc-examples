package com.dinhhuy258.distributedlock.controllers

import com.dinhhuy258.distributedlock.services.DistributedLock
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class DemoLockController(private val distributedLock: DistributedLock) {

    @GetMapping("/lock-demo")
    fun getLockDemo(): String {
        distributedLock.lock()

        Thread.sleep(5000)

        distributedLock.unlock()

        return "Done";
    }
}
