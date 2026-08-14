import Link from "next/link";

export default function ProductCard({ product }) {
  return (
    <div className="product-card">
      <h2>{product.name}</h2>

      <p className="category">
        Category: {product.category}
      </p>

      <p className="price">
        ₹{product.price}
      </p>

      <Link
        className="details-button"
        href={`/products/${product.id}`}
      >
        View Details
      </Link>
    </div>
  );
}
