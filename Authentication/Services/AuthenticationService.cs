using System;
using System.Security.Claims;
using System.Linq;
using Microsoft.IdentityModel.Tokens;
using System.IdentityModel.Tokens.Jwt;
using Microsoft.EntityFrameworkCore;
using Comies;
using System.Security.Principal;

namespace Comies.Auth
{
    public class AuthenticationService
    {
        private SigningConfigurations SigningConfigurations { get; set; }
        private AuthenticationConfiguration TokenConfigurations { get; set; }
        private ComiesContext Context { get; set; }

        public DateTime ExpirationDate => DateTime.Now.AddSeconds(TokenConfigurations.Seconds);

        public AuthenticationService(SigningConfigurations signingConfigurations, AuthenticationConfiguration tokenConfigurations, ComiesContext context)
        {
            SigningConfigurations = signingConfigurations;
            TokenConfigurations = tokenConfigurations;
            Context = context;
        }

        public Operator GetOperator(AuthenticationParameters applicant) 
        {
            var applicantData = ValidateAndExtractNicknameAndStore(applicant);
            var storeId = (from s in Context.Stores where s.Active == true && s.CompanyNickname == applicantData[1] select s.Id).FirstOrDefault();
            if (storeId == Guid.Empty) throw new ComiesUnauthorizedException("Ops, não foi possível encontrar seu domínio. Contate o administrador de sua empresa ;)");
            var foundOperator = (from o in Context.Operators where o.Active == true && o.Nickname == applicantData[0] && o.StoreId == storeId select o).FirstOrDefault();
            if (foundOperator == null) throw new ComiesUnauthorizedException("Ops, acho que não nos conhecemos, pois não foi possível encontrar você com esse apelido.");
            if (foundOperator.Password.Equals(applicant.Password)) return foundOperator;
            else throw new ComiesUnauthorizedException("Ops! Senha incorreta. Você digitou tudo certinho?");
        }

        /// <summary>
        /// Validates if applicant has the operator nickname and the store nickname, and if it has a valid password
        /// </summary>
        /// <param name="applicant">The object containing the information to be validated</param>
        /// <returns>A string array containg [0] the nickname and [1] the store </returns>
        private static string[] ValidateAndExtractNicknameAndStore(AuthenticationParameters applicant)
        {
            if (applicant == null) throw new ComiesArgumentException(message: "Ops! Você precisa nos passar alguma informação que a gente te reconheça.", paramName: nameof(applicant));
            if (string.IsNullOrEmpty(applicant.Nickname)) throw new ComiesArgumentException(message: "Ops! Precisamos do seu apelido para te identificar", paramName: nameof(applicant));
            if (!applicant.Nickname.Contains('@') && applicant.Nickname.Length < 7) throw new ComiesArgumentException(message: "Ops! Seu apelido deve ser composto de apelido@nome_da_empresa", paramName: nameof(applicant));
            if (string.IsNullOrEmpty(applicant.Password)) throw new ComiesArgumentException(message: "Ops! Precisamos da sua senha para que você possa se autenticar", paramName: nameof(applicant));
            return applicant.Nickname.Split('@');
        }

        /// <summary>
        /// Gets the access token for the operator
        /// </summary>
        /// <param name="oper">The operator to receive the token</param>
        /// <returns>The token fetched</returns>
        public string GetToken(Operator oper, bool keepConnected = false)
        {
            return GetToken(GetClaimsIdentity(oper), DateTime.Now, ExpirationDate.AddDays(keepConnected ? 1 : 0));
        }

        private ClaimsIdentity GetClaimsIdentity(Operator op)
        {
            var profile = (from p in Context.Profiles where p.Id == op.ProfileId select p).FirstOrDefault();
            return new ClaimsIdentity(
                new GenericIdentity(op.Name, "Name"),
                new[]
                {
                    new Claim("OperatorId", op.Id.ToString()), new Claim("StoreId", op.StoreId.ToString()),
                    new Claim("Nickname", op.Nickname), new Claim("ProfileId", profile == null ? "" :  profile.Id.ToString() )                  
                }
            );
        }

        /// <summary>
        /// Creates the token for this access
        /// </summary>
        /// <param name="identity">the Claim object</param>
        /// <param name="creationDate">the creation date of this token</param>
        /// <param name="expirationDate">the expiration date of this object </param>
        /// <returns>The fetched token</returns>
        private string GetToken(ClaimsIdentity identity, DateTime creationDate, DateTime expirationDate)
        {
            var handler = new JwtSecurityTokenHandler();
            var token = handler.CreateToken(new SecurityTokenDescriptor {
                Issuer = TokenConfigurations.Issuer,
                Audience = TokenConfigurations.Audience,
                SigningCredentials = SigningConfigurations.SigningCredentials,
                Subject = identity,
                NotBefore = creationDate,
                Expires = expirationDate
            });
            return handler.WriteToken(token);
        }
    }
}