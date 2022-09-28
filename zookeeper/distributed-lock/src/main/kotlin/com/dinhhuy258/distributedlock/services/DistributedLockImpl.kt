package com.dinhhuy258.distributedlock.services

import org.apache.curator.framework.CuratorFramework
import org.apache.curator.framework.recipes.locks.InterProcessSemaphoreMutex
import org.springframework.stereotype.Service

@Service
class DistributedLockImpl(curatorFramework: CuratorFramework) : DistributedLock {
    private val lockPath = "/lock"

    private val mutex: InterProcessSemaphoreMutex = InterProcessSemaphoreMutex(curatorFramework, lockPath)

    override fun lock() {
        mutex.acquire()
    }

    override fun unlock() {
        mutex.release()
    }
}
