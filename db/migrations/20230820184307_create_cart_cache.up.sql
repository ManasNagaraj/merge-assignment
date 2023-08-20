CREATE TABLE `cart_caches` (
  `id` uuid NOT NULL DEFAULT uuid(),
  `item_id` varchar(255) NOT NULL,
  `user_id` varchar(255) NOT NULL,
  `quantity` int(11) DEFAULT NULL,
  `expiry` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;