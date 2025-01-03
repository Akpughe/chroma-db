use crate::execution::operator::Operator;
use async_trait::async_trait;
use chroma_distance::normalize;

#[derive(Debug)]
pub struct NormalizeVectorOperator {}

pub struct NormalizeVectorOperatorInput {
    pub vectors: Vec<Vec<f32>>,
}

pub struct NormalizeVectorOperatorOutput {
    pub _normalized_vectors: Vec<Vec<f32>>,
}

#[async_trait]
impl Operator<NormalizeVectorOperatorInput, NormalizeVectorOperatorOutput>
    for NormalizeVectorOperator
{
    type Error = ();

    fn get_name(&self) -> &'static str {
        "NormalizeVectorOperator"
    }

    async fn run(
        &self,
        input: &NormalizeVectorOperatorInput,
    ) -> Result<NormalizeVectorOperatorOutput, Self::Error> {
        // TODO: this should not have to reallocate the vectors. We can optimize this later.
        let mut normalized_vectors = Vec::with_capacity(input.vectors.len());
        for vector in &input.vectors {
            let normalized_vector = normalize(vector);
            normalized_vectors.push(normalized_vector);
        }
        Ok(NormalizeVectorOperatorOutput {
            _normalized_vectors: normalized_vectors,
        })
    }
}

#[cfg(test)]
mod tests {
    use super::*;

    const COMPARE_EPS: f32 = 1e-9;
    fn float_eps_eq(a: &[f32], b: &[f32]) -> bool {
        a.iter()
            .zip(b.iter())
            .all(|(a, b)| (a - b).abs() < COMPARE_EPS)
    }

    #[tokio::test]
    async fn test_normalize_vector() {
        let operator = NormalizeVectorOperator {};
        let input = NormalizeVectorOperatorInput {
            vectors: vec![
                vec![1.0, 2.0, 3.0],
                vec![4.0, 5.0, 6.0],
                vec![7.0, 8.0, 9.0],
            ],
        };

        let output = operator.run(&input).await.unwrap();
        let expected_output = NormalizeVectorOperatorOutput {
            _normalized_vectors: vec![
                vec![0.26726124, 0.5345225, 0.8017837],
                vec![0.45584232, 0.5698029, 0.68376344],
                vec![0.5025707, 0.5743665, 0.64616233],
            ],
        };

        for (a, b) in output
            ._normalized_vectors
            .iter()
            .zip(expected_output._normalized_vectors.iter())
        {
            assert!(float_eps_eq(a, b), "{:?} != {:?}", a, b);
        }
    }
}
