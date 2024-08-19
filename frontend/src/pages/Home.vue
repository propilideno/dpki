<template>
  <q-page class="q-pa-md" padding>
    <q-card>
      <q-card-section class="q-pb-none">
        <div class="text-h5">
          DPKI - Frequently Asked Questions
        </div>
      </q-card-section>
      <q-card-section>
        <q-list bordered class="rounded-borders q-pa-none">
          <q-expansion-item
            icon="help"
            label="What is a Decentralized Public Key Infrastructure (DPKI)?"
            header-class="text-subtitle1"
          >
            <q-card>
              <q-card-section>
                <p>
                  A Decentralized Public Key Infrastructure (DPKI) is a system that manages public keys without relying on a central authority.
                    Unlike traditional PKIs, which depend on a central Certificate Authority (CA) to issue and validate digital certificates,
                    DPKI operates on a decentralized network, often utilizing blockchain technology.
                    This approach enhances security and trust by eliminating single points of failure and reducing the risk of central authority compromise.
                    In DPKI, users control their own keys and identities, making it more resilient, transparent, and suitable for decentralized applications.
                </p>
              </q-card-section>
            </q-card>
          </q-expansion-item>
    
          <q-expansion-item
            icon="work"
            label="How does DPKI work?"
            header-class="text-subtitle1"
          >
            <q-card>
              <q-card-section>
                <p>
                  Decentralized Public Key Infrastructure (DPKI) operates by distributing the management of public keys across a decentralized network, rather than relying on a central authority like in traditional PKI systems. Here’s how DPKI works:
                </p>
                <h6>1. Key Generation and Control:</h6>
                <p>
                  In DPKI, users generate their own public and private key pairs. They maintain control over their private keys, which are never shared with any third party.
                  The public key, which is needed to verify digital signatures or encrypt data, can be published to the decentralized network.
                </p>
                <h6>2. Decentralized Verification:</h6>
                <p>
                  Unlike traditional PKI where a Certificate Authority (CA) verifies and vouches for the authenticity of a public key, DPKI relies on the network itself for verification.
                  This is often achieved using blockchain or other distributed ledger technologies. The blockchain acts as an immutable and transparent ledger, where public keys can be recorded and linked to user identities or digital signatures.
                </p>
                <h6>
                  3. Identity Management:
                </h6>
                <p>
                  DPKI may use Decentralized Identifiers (DIDs), which are self-sovereign and allow users to create and control their own digital identities. These identities are not dependent on any central authority.
                  The DID is associated with the user’s public key, and this association is recorded on the blockchain.
                </p>
                <h6>
                  4. Trust and Consensus:
                </h6>
                <p>
                  In DPKI, trust is established through a consensus mechanism that ensures the integrity and authenticity of the public keys and identities recorded on the network.
                  The decentralized nature of the network means that multiple nodes participate in validating transactions, making it highly resistant to tampering or fraud.
                </p>
                <h6>
                  5. Revocation and Updates:
                </h6>
                <p>
                  Revocation or updating of keys is also managed through the decentralized network. If a key needs to be revoked or replaced, this information is propagated across the network and recorded on the blockchain.
                  This ensures that any changes to the key status are transparent and verifiable by anyone on the network.
                </p>
                <h6>
                  6. Security and Privacy:
                </h6>
                <p>
                  DPKI enhances security by removing the single points of failure inherent in centralized systems. Since there is no central authority that can be compromised, the risk of large-scale breaches is reduced.
                  Additionally, privacy is enhanced because users control their own identities and keys, without needing to rely on or trust third parties.
                </p>
                <h6>
                  Summary:
                </h6>
                <p>
                  DPKI decentralizes the trust model by distributing the management and verification of public keys across a blockchain or distributed ledger, empowering users to control their own identities and enhancing the security and resilience of the system.
                </p>
              </q-card-section>
            </q-card>
          </q-expansion-item>
          <q-expansion-item
            icon="info"
            label="PKIs vs. DPKIs"
            header-class="text-subtitle1"
          >
            <q-card>
              <q-card-section>
                <p>
                  The main difference between a PKI (Public Key Infrastructure) and a DPKI (Decentralized Public Key Infrastructure) lies in how trust and key management are handled. In a traditional PKI, trust is centralized in a Certificate Authority (CA), which is responsible for issuing, validating, and revoking digital certificates. These certificates link a public key to an identity, and users rely on the CA to ensure the authenticity of these associations. This means the security of a PKI heavily depends on the integrity and security of the CA. If the CA is compromised, the entire system can become vulnerable, exposing users to risks such as man-in-the-middle attacks.
                </p>
                <p>
                  In contrast, a DPKI distributes key management and trust across a decentralized network, often using blockchain technology. Instead of relying on a single central entity, users trust the network's consensus to verify the authenticity of public keys and their associations with identities. This eliminates the single point of failure present in centralized systems, making DPKI more resilient to attacks. Additionally, in a DPKI, users have full control over their digital identities, allowing them to manage their own keys without depending on a central authority.
                </p>
                <p>
                  This decentralization also impacts cost and scalability. While traditional PKI can be expensive and complex to manage, especially in large organizations, DPKI distributes cost and responsibility across the network, potentially reducing overall expenses. However, DPKI can face scalability and latency challenges depending on the underlying technology, such as blockchain, which may introduce new complexities. In summary, while PKI centralizes trust and is well-established, DPKI offers a more decentralized and resilient approach, though it comes with its own set of challenges.
                </p>
              </q-card-section>
            </q-card>
          </q-expansion-item>
        </q-list>
      </q-card-section>
    </q-card>

  </q-page>
</template>

<script>

export default {
  name: 'HomePage',
}
</script>

<script setup>

</script>

<style lang="scss">
h6 {
  margin-top: 0;
  margin-bottom: 0;
}
</style>