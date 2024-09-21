package com.soal.model;

import lombok.Data;
import jakarta.persistence.*;
import java.util.UUID;

@Entity
@Data
public class User {
    @Id
    @GeneratedValue
    private UUID id;

    @Column(unique = true)
    private String phoneNumber;

    private String firstName;
    private String lastName;
    private String address;
    private String pin; // Store hashed PIN
    private Integer balance = 0;

    private Boolean isActive = true;
}

