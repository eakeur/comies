using Microsoft.EntityFrameworkCore.Metadata.Builders;
using Microsoft.EntityFrameworkCore;
using Comies;

namespace Comies.ModelsSettings
{   
    public class ProfileConfiguration : IEntityTypeConfiguration<Profile>{
        
        public void Configure(EntityTypeBuilder<Profile> builder){
            builder.HasKey(p => p.Id);
            builder.Property(p => p.ProfileName).IsRequired(true);

            builder.HasIndex(i => new { i.ProfileName, i.StoreId }).IsUnique();
        }
    }
}