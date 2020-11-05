class Property {
  final int bedrooms;
  final int bathrooms;
  final int garages;
  final bool pool;
  final String address;
  final String size;
  final String type;

  Property(this.bedrooms, this.bathrooms, this.garages, this.pool, this.address,
      this.size, this.type);

  Map<String, dynamic> toJson() {
    return {
      "Bedrooms": bedrooms,
      "Bathrooms": bathrooms,
      "Garages": garages,
      "Pool": pool,
      "Address": address,
      "Size": size,
      "Type": type
    };
  }
}
