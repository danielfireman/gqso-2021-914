/**
 * @param {import('express').Request} req - Request
 * @param {import('express').Response} res - Response.
 */
function raizQuadrada(req, res) {
    try {
        const op = Math.sqrt(parseFloat(req.params.op));
        if (typeof op !== Number) {
            throw new Error('invalid param');
        }
        return res.send('' + op)
    } catch (err) {
        return res.status(400).send(`Parâmetro Inválido: ${req.params.op}`);
    }
}

module.exports = { raizQuadrada }