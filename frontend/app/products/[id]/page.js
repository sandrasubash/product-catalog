"use client";

import { useEffect, useState } from "react";
import Link from "next/link";
import { getProduct } from "../../../lib/api";

export default function ProductDetails({ params }) {
  const [product, setProduct] = useState(null);
  const [loading, setLoading] = useState(true);
  const [notFound, setNotFound] = useState(false);
  const [error, setError] = useState("");

  useEffect(() => {
    async function loadProduct() {
      try {
        setLoading(true);
        setError("");
        setNotFound(false);

        const { id } = await params;

        const data = await getProduct(id);

        setProduct(data);
      } catch (error) {
        console.error(error);

        if (error.message === "Product not found") {
          setNotFound(true);
        } else {
          setError(
            "Something went wrong. Please try again."
          );
        }
      } finally {
        setLoading(false);
      }
    }

    loadProduct();
  }, [params]);

  if (loading) {
    return (
      <main className="container">
        <p className="status-message">
          Loading...
        </p>
      </main>
    );
  }

  if (notFound) {
    return (
      <main className="container">
        <h1>Product not found</h1>

        <Link href="/" className="back-link">
          Back to Products
        </Link>
      </main>
    );
  }

  if (error) {
    return (
      <main className="container">
        <p className="error-message">
          {error}
        </p>

        <Link href="/" className="back-link">
          Back to Products
        </Link>
      </main>
    );
  }

  return (
    <main className="container">
      <div className="product-detail">
        <h1>{product.name}</h1>

        <p>
          <strong>Category:</strong>{" "}
          {product.category}
        </p>

        <p className="price">
          ₹{product.price}
        </p>

        <Link href="/" className="back-link">
          ← Back to Products
        </Link>
      </div>
    </main>
  );
}
