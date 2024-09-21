package com.soal.repository;

import com.soal.model.Transaction;
import org.springframework.data.jpa.repository.JpaRepository;
import java.util.List;
import java.util.UUID;

public interface TransactionRepository extends JpaRepository<Transaction, UUID> {
    
    // Fetch all transactions by user id
    List<Transaction> findByUserId(UUID userId);
}
