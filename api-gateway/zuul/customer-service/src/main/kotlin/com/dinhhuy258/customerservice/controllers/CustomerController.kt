package com.dinhhuy258.customerservice.controllers

import com.dinhhuy258.customerservice.controllers.dtos.CustomerDTO
import org.springframework.web.bind.annotation.GetMapping
import org.springframework.web.bind.annotation.RestController

@RestController
class CustomerController {
    @GetMapping("/customers")
    fun getCustomers(): List<CustomerDTO> {
        return listOf(CustomerDTO(1, "Huy"), CustomerDTO(2, "Thuy"));
    }
}
