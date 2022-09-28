package com.dinhhuy258.distributedlock.services

interface DistributedLock {
    fun lock()

    fun unlock()
}
