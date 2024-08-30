-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Aug 30, 2024 at 12:57 PM
-- Server version: 9.0.1
-- PHP Version: 8.2.22

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_xyz_finance`
--

-- --------------------------------------------------------

--
-- Table structure for table `access_apps`
--

CREATE TABLE `access_apps` (
  `id` bigint UNSIGNED NOT NULL,
  `domain` varchar(191) DEFAULT NULL,
  `api_key` longtext,
  `updated_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `access_apps`
--

INSERT INTO `access_apps` (`id`, `domain`, `api_key`, `updated_at`, `created_at`) VALUES
(1, 'localhost:8080', '1n1_4p1_Key', NULL, NULL);

-- --------------------------------------------------------

--
-- Table structure for table `customers`
--

CREATE TABLE `customers` (
  `id` bigint UNSIGNED NOT NULL,
  `nik` varchar(255) DEFAULT NULL,
  `full_name` varchar(255) DEFAULT NULL,
  `legal_name` varchar(255) DEFAULT NULL,
  `place_of_birth` varchar(255) DEFAULT NULL,
  `date_of_birth` date DEFAULT NULL,
  `salary` double DEFAULT NULL,
  `id_card_photo` varchar(255) DEFAULT NULL,
  `selfie_photo` varchar(255) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `customers`
--

INSERT INTO `customers` (`id`, `nik`, `full_name`, `legal_name`, `place_of_birth`, `date_of_birth`, `salary`, `id_card_photo`, `selfie_photo`, `updated_at`, `created_at`) VALUES
(1, '3511061511030001', 'Budi', 'Budi', 'Bondowoso', '2002-08-29', 5000000, 'uploads/KTP/Screenshot 2024-08-21 at 16.32.44.png', 'uploads/Selfie/Screenshot 2024-08-21 at 16.33.00.png', '2024-08-30 19:46:23.006', '2024-08-30 19:46:23.006'),
(2, '3511061511030002', 'Annisa', 'Annisa', 'Bondowoso', '2002-08-29', 10000000, 'uploads/KTP/Screenshot 2024-08-21 at 16.32.44.png', 'uploads/Selfie/Screenshot 2024-08-21 at 16.33.00.png', '2024-08-30 19:56:16.399', '2024-08-30 19:56:16.399');

-- --------------------------------------------------------

--
-- Table structure for table `customer_limits`
--

CREATE TABLE `customer_limits` (
  `id` bigint UNSIGNED NOT NULL,
  `customer_id` bigint UNSIGNED NOT NULL,
  `term` bigint DEFAULT NULL,
  `amount` decimal(10,2) NOT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data for table `customer_limits`
--

INSERT INTO `customer_limits` (`id`, `customer_id`, `term`, `amount`, `updated_at`, `created_at`) VALUES
(1, 1, 1, 100000.00, '2024-08-30 19:54:07.820', '2024-08-30 19:54:07.820'),
(2, 1, 2, 200000.00, '2024-08-30 19:54:16.118', '2024-08-30 19:54:16.118'),
(3, 1, 3, 500000.00, '2024-08-30 19:54:28.757', '2024-08-30 19:54:28.757'),
(4, 1, 6, 700000.00, '2024-08-30 19:55:23.499', '2024-08-30 19:55:23.499'),
(6, 2, 1, 1000000.00, '2024-08-30 19:56:27.213', '2024-08-30 19:56:27.213'),
(7, 2, 2, 1200000.00, '2024-08-30 19:56:39.198', '2024-08-30 19:56:39.198'),
(8, 2, 3, 1500000.00, '2024-08-30 19:56:48.102', '2024-08-30 19:56:48.102'),
(9, 2, 6, 2000000.00, '2024-08-30 19:56:56.648', '2024-08-30 19:56:56.648');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint UNSIGNED NOT NULL,
  `customer_id` bigint UNSIGNED NOT NULL,
  `customer_limit_id` bigint UNSIGNED NOT NULL,
  `contract_number` longtext,
  `otr` double DEFAULT NULL,
  `admin_fee` double DEFAULT NULL,
  `installment_amount` double DEFAULT NULL,
  `interest_amount` double DEFAULT NULL,
  `asset_name` longtext,
  `status` longtext,
  `term` bigint DEFAULT NULL,
  `billing_date` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `created_at` datetime(3) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `access_apps`
--
ALTER TABLE `access_apps`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `uni_access_apps_domain` (`domain`);

--
-- Indexes for table `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `customer_limits`
--
ALTER TABLE `customer_limits`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_customers_customer_limit` (`customer_id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_customers_transaction` (`customer_id`),
  ADD KEY `fk_transactions_customer_limit` (`customer_limit_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `access_apps`
--
ALTER TABLE `access_apps`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `customers`
--
ALTER TABLE `customers`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `customer_limits`
--
ALTER TABLE `customer_limits`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint UNSIGNED NOT NULL AUTO_INCREMENT;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `customer_limits`
--
ALTER TABLE `customer_limits`
  ADD CONSTRAINT `fk_customers_customer_limit` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `fk_customers_transaction` FOREIGN KEY (`customer_id`) REFERENCES `customers` (`id`),
  ADD CONSTRAINT `fk_transactions_customer_limit` FOREIGN KEY (`customer_limit_id`) REFERENCES `customer_limits` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
