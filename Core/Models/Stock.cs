using Comies;
using System.Collections.Generic;
using System;
using System.ComponentModel.DataAnnotations.Schema;
namespace Comies
{
    public class Stock : StoreOwnedEntity
    {
        public Guid ProductId { get; set; }
		public DateTime Date { get; set; }
		public double Minimum { get; set; }
		public double Maximum { get; set; }
		public double Actual { get; set; }
		public string Location { get; set; }
		public Unity StockUnity { get; set; }

		[System.Text.Json.Serialization.JsonIgnore]
		public Product Product { get; set; }
        public virtual IList<StockMovement> Movements { get; set; }

    }
}