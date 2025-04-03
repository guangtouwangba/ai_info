# Review, Refine, Repeat: Understanding Iterative Decoding of AI Agents with Dynamic Evaluation and Selection

**URL**: http://arxiv.org/abs/2504.01931v1

## 原始摘要

While AI agents have shown remarkable performance at various tasks, they
still struggle with complex multi-modal applications, structured generation and
strategic planning. Improvements via standard fine-tuning is often impractical,
as solving agentic tasks usually relies on black box API access without control
over model parameters. Inference-time methods such as Best-of-N (BON) sampling
offer a simple yet effective alternative to improve performance. However, BON
lacks iterative feedback integration mechanism. Hence, we propose Iterative
Agent Decoding (IAD) which combines iterative refinement with dynamic candidate
evaluation and selection guided by a verifier. IAD differs in how feedback is
designed and integrated, specifically optimized to extract maximal signal from
reward scores. We conduct a detailed comparison of baselines across key metrics
on Sketch2Code, Text2SQL, and Webshop where IAD consistently outperforms
baselines, achieving 3--6% absolute gains on Sketch2Code and Text2SQL (with and
without LLM judges) and 8--10% gains on Webshop across multiple metrics. To
better understand the source of IAD's gains, we perform controlled experiments
to disentangle the effect of adaptive feedback from stochastic sampling, and
find that IAD's improvements are primarily driven by verifier-guided
refinement, not merely sampling diversity. We also show that both IAD and BON
exhibit inference-time scaling with increased compute when guided by an optimal
verifier. Our analysis highlights the critical role of verifier quality in
effective inference-time optimization and examines the impact of noisy and
sparse rewards on scaling behavior. Together, these findings offer key insights
into the trade-offs and principles of effective inference-time optimization.


## AI 摘要

当前AI代理在处理多模态应用、结构化生成和策略规划等复杂任务时仍有不足。由于通常依赖黑盒API而无法调整模型参数，传统微调方法效果有限。研究者提出了一种迭代代理解码（IAD）方法，通过动态候选评估和验证器引导的迭代优化来提升性能。在Sketch2Code、Text2SQL和Webshop等任务中，IAD比基线方法绝对提升3-10%。实验表明，IAD的改进主要来自验证器引导的优化而非采样多样性。研究还揭示了验证器质量对推理时优化的关键作用，以及噪声/稀疏奖励对扩展行为的影响。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-03T03:13:18Z
- **目录日期**: 2025-04-03
