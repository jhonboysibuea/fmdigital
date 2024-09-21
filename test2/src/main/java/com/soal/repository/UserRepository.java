package com.soal.repository;

import org.springframework.data.jpa.repository.JpaRepository;
import java.util.Optional;
import java.util.UUID;
import com.soal.model.*;

public interface UserRepository extends JpaRepository<User, UUID> {
    
    // Find user by phone number (unique key)
    Optional<User> findByPhoneNumber(String phoneNumber);

    // Check if phone number already exists
    boolean existsByPhoneNumber(String phoneNumber);
}