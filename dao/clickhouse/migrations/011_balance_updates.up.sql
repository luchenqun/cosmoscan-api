create table balance_updates
(
    bau_id         FixedString(40),
    bau_address    FixedString(45),
    bau_balance    Decimal(32, 8),
    bau_stake      Decimal(32, 8),
    bau_unbonding  Decimal(32, 8),
    bau_created_at DateTime
) ENGINE ReplacingMergeTree()
      PARTITION BY toYYYYMMDD(bau_created_at)
      ORDER BY (bau_id);

