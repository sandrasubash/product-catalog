"use client";

export default function CategoryFilter({
  category,
  onCategoryChange,
}) {
  return (
    <div className="filter-container">
      <label htmlFor="category">
        Filter by Category:
      </label>

      <select
        id="category"
        value={category}
        onChange={(event) =>
          onCategoryChange(event.target.value)
        }
      >
        <option value="">All Products</option>
        <option value="Clothing">Clothing</option>
        <option value="Kitchen">Kitchen</option>
        <option value="Accessories">Accessories</option>
        <option value="Footwear">Footwear</option>
      </select>
    </div>
  );
}
