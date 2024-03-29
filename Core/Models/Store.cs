using System.Collections.Generic;
using System.ComponentModel.DataAnnotations;
using System;
using System.ComponentModel.DataAnnotations.Schema;
namespace Comies
{
    public class Store
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)] 
        public Guid Id { get; set; }
        public string Name { get; set; }

        [MaxLength(20)]
        public string CompanyNickname { get; set; }

        [MaxLength(200)]
        public string CompanyName { get; set; }
        public string Document { get; set; }
        public Guid? StoreId { get; set; }
        public Store Parent { get; set; }
        public bool Active { get; set; }
        public DateTime CreationDate { get; set; }

        [MaxLength(200)]
        public string ContactName { get; set; }
        public virtual DateTime MemberSince { get; set; }  
        public virtual IList<Store> Stores { get; set; }
        public virtual IList<Product> Products { get; set; }
        public virtual IList<ProductCategory> ProductCategories { get; set; }
        public virtual IList<Operator> Operators { get; set; }
        public virtual IList<Order> Orders { get; set; }
        public virtual IList<Profile> Profiles { get; set; }
        public virtual IList<Stock> Stocks { get; set; }
        public virtual IList<StockMovement> StockMovements { get; set; }
        public virtual IList<StoreProperty> StoreProperties { get; set; }

    }
}