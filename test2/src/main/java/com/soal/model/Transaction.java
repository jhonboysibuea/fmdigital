package com.soal.model;


import lombok.Data;
import jakarta.persistence.*;

import jakarta.persistence.Entity;

import java.util.UUID;
import java.time.LocalDateTime;

@Entity
@Data
public class Transaction {
    @Id
    @GeneratedValue
    private UUID id;

    @ManyToOne
    @JoinColumn(name = "user_id", nullable = false)
    private User user;

    private String type; // "CREDIT", "DEBIT"
    private Integer amount;
    private String remarks;
    private Integer balanceBefore;
    private Integer balanceAfter;
    private LocalDateTime createdDate = LocalDateTime.now();
}