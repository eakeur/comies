using System;
using System.ComponentModel.DataAnnotations.Schema;

namespace Comies
{
    public class OrderItem : StoreOwnedEntity
    {
        
        public Guid OrderId { get; set; }
        public int Group { get; set; }
        public double Quantity { get; set; }
        public bool Done { get; set; }
        public decimal Discount { get; set; }
        public decimal Price { get; set; }
        public decimal FinalPrice { get; set; }
        public Guid ProductId { get; set; }
        public virtual Product Product { get; set; }
        public virtual Order Order { get; set; }
    }
}