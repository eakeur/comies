using System.ComponentModel.DataAnnotations.Schema;
using System.ComponentModel.DataAnnotations;
using System;
using Newtonsoft.Json.Linq;

namespace Comies {
    public abstract class StoreOwnedEntity : Entity
    {
        
        [Required(ErrorMessage="É preciso especificar uma loja à qual este item pertence")]
        public Guid StoreId { get; set; }

        [System.Text.Json.Serialization.JsonIgnore]
        public virtual Store Store { get; set; }
    }

    public abstract class Entity
    {
        [DatabaseGenerated(DatabaseGeneratedOption.Identity)]
        public Guid Id { get; set; }
        public Entity() { CreationDate = DateTime.Now;}
        public bool Active { get; set; }
        public DateTime CreationDate { get; set; }

        
    }

    


}