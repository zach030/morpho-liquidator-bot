pragma solidity ^0.8.13;
import "@src/Interface.sol";

interface IBot {
    struct MorphoLiquidateData {
        address collateralToken;
        address loanToken;
        uint256 seized;
        address pair;
        bytes swapData;
    }
    function withdrawETH(uint256 amount) external;
    function withdrawERC20(address token, uint256 amount) external;
    function approveERC20(address token, address to, uint256 amount) external;
    function morphoLiquidate(bytes32 id, address borrower, uint256 seizedAssets, address pair, bytes calldata swapData) external payable;
}