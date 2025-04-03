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

当前AI代理在多模态应用、结构化生成和策略规划等复杂任务中仍有不足。由于通常依赖黑盒API且无法调整模型参数，传统微调方法效果有限。研究提出迭代代理解码（IAD）方法，通过验证器指导的动态候选评估和选择实现迭代优化。相比最佳N采样（BON），IAD能更有效利用奖励信号。在Sketch2Code、Text2SQL和Webshop任务中，IAD性能显著优于基线（提升3-10%）。实验表明其优势主要来自验证器引导的优化而非采样多样性。研究还揭示了验证器质量和计算资源对推理时优化的重要影响。

## 元数据

- **来源**: ArXiv
- **类型**: 论文
- **保存时间**: 2025-04-03T05:01:04Z
- **目录日期**: 2025-04-03
